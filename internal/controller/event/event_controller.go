package event

import (
	"KUNoti/internal/request/eventrequest"
	eventservice "KUNoti/service/event"
	"KUNoti/service/firebaseService"
	"KUNoti/service/s3service"
	"KUNoti/sqlc"
	"context"
	"encoding/json"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EventController struct {
	es *eventservice.EventService
	s3 *s3service.S3Service
	fb firebaseService.FireBaseService
}

func (e EventController) CreateEvent(ctx *gin.Context) {
	var createEventRequest eventrequest.CreateEventRequest
	err := ctx.ShouldBind(&createEventRequest)
	if err != nil {
		log.Println(err.Error())
		log.Printf("Error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if createEventRequest.ImageFile != nil {
		imageURL, err := e.s3.Upload(s3service.EventImageFolder, createEventRequest.ImageFile)
		if err != nil {
			log.Println("Error saving image to S3:", err)
			ctx.JSON(http.StatusInternalServerError, "Error saving image")
			return
		}
		createEventRequest.Image = imageURL
	}

	event, err := e.es.Create(ctx, createEventRequest)
	if err != nil {
		log.Println(err)
	}

	eventData, err := json.Marshal(event)
	if err != nil {
		log.Printf("Error serializing event data: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing event data"})
		return
	}

	newCtx := ctx.Copy()
	go func() {
		tokens, err := e.es.FindTokensByTagName(newCtx, createEventRequest.Tag)
		if err != nil {
			log.Printf("Error finding tokens: %v", err)
			return
		}
		if len(tokens) > 0 {
			err := e.fb.SendMulticastWithData(newCtx, tokens, "New Event: "+event.Title, "Check out this new event happening soon!", eventData)
			if err != nil {
				log.Printf("Error sending notifications: %v", err)
			}
		}
	}()

	ctx.JSON(201, createEventRequest)
}

func (e EventController) UpdateEvent(ctx *gin.Context) {
	var updateEventRequest eventrequest.UpdateEventRequest
	err := ctx.Bind(&updateEventRequest)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	event, err := e.es.Update(ctx, updateEventRequest)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, event)
}

func (e EventController) DeleteEvent(ctx *gin.Context) {
	var deleteEventRequest eventrequest.DeleteEventRequest
	err := ctx.BindJSON(&deleteEventRequest)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	id, err := e.es.Delete(ctx, deleteEventRequest)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, "delete event ID : "+id)
}

func (e EventController) FinderEvent(ctx *gin.Context) {
	var finderEventRequest eventrequest.FinderEventRequest
	if queryParam, ok := ctx.GetQuery("keyword"); ok {
		finderEventRequest.Keyword = queryParam
	}

	events, err := e.es.Finder(ctx, finderEventRequest)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200, events)
}

func (e EventController) Events(ctx *gin.Context) {
	events, err := e.es.FindAll(ctx)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, events)
}

func (e EventController) FollowEvent(ctx *gin.Context) {
	var followEventRequest eventrequest.FollowEventRequest
	err := ctx.BindJSON(&followEventRequest)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	followE, err := e.es.Follow(ctx, followEventRequest)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	str := "User ID: " + strconv.Itoa(followE.UserID) + " following Event ID: " + strconv.Itoa(followE.EventID)
	ctx.JSON(200, str)
}

func (e EventController) UnFollowEvent(ctx *gin.Context) {
	var unfollowEventRequest eventrequest.UnfollowEventRequest
	err := ctx.BindJSON(&unfollowEventRequest)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	id, err := e.es.Unfollow(ctx, unfollowEventRequest)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200, "unfollow event ID : "+id)
}

func (e EventController) FollowEvents(ctx *gin.Context) {
	var findFollowEventRequest eventrequest.FindFollowEventRequest
	err := ctx.BindJSON(&findFollowEventRequest)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	events, err := e.es.FindFollowEvent(ctx, findFollowEventRequest)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200, events)
}

func (e EventController) FindEventCreatedByMe(ctx *gin.Context) {
	var findEventCreatedByMeRequest eventrequest.FindEventCreatedByMeRequest
	err := ctx.BindJSON(&findEventCreatedByMeRequest)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	events, err := e.es.FindEventCreatedByMe(ctx, findEventCreatedByMeRequest)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200, events)
}

func (e EventController) FindTagByToken(ctx *gin.Context) {
	var request eventrequest.FindTagByToken
	err := ctx.BindJSON(&request)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	tag, err := e.es.FindTagByToken(ctx, request)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200, tag)
}

func (e EventController) FollowTag(ctx *gin.Context) {
	var followTagRequest eventrequest.FollowTagRequest
	err := ctx.BindJSON(&followTagRequest)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	followT, err := e.es.FollowTag(ctx, followTagRequest)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200, "follow "+followT.Tag)
}

func (e EventController) UnFollowTag(ctx *gin.Context) {
	var unFollowTagRequest eventrequest.UnFollowTagRequest
	err := ctx.BindJSON(&unFollowTagRequest)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	tag, err := e.es.UnfollowTag(ctx, unFollowTagRequest)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200, "unfollow "+tag)
}

func (e EventController) RegisEvent(ctx *gin.Context) {
	var regisEventRequest eventrequest.RegisEventRequest
	err := ctx.BindJSON(&regisEventRequest)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	regisE, err := e.es.RegisEvent(ctx, regisEventRequest)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200, regisE)
}

func (e EventController) RegisEvents(ctx *gin.Context) {
	var regisEventRequest eventrequest.RegisEventRequest
	err := ctx.BindJSON(&regisEventRequest)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	regisEvents, err := e.es.FindRegisEvent(ctx, regisEventRequest)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200, regisEvents)
}

func (e EventController) Notification(ctx context.Context, tokens []string, event sqlc.Event, eventData []byte) {
	go func() {
		err := e.fb.SendMulticastWithData(ctx, tokens, "New Event: "+event.Title, "Check out this new event happening soon!", eventData)
		if err != nil {
			log.Printf("Error sending notifications: %v", err)
		}
	}()
}

type getNotification struct {
	Token string `json:"token"`
}

func (e EventController) Notifications(ctx *gin.Context) {
	var notificationsRequest getNotification
	err := ctx.BindJSON(&notificationsRequest)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	notis, err := e.fb.Notifications(ctx, notificationsRequest.Token)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200, notis)
}

func (e EventController) InitEndpoints(r *gin.RouterGroup) {
	eventGroup := r.Group("/event")
	eventGroup.POST("/create", e.CreateEvent)
	eventGroup.PUT("/update", e.UpdateEvent)
	eventGroup.DELETE("/delete", e.DeleteEvent)
	eventGroup.GET("", e.FinderEvent)
	eventGroup.GET("/events", e.Events)
	eventGroup.POST("/follow", e.FollowEvent)
	eventGroup.DELETE("/unfollow", e.UnFollowEvent)
	eventGroup.GET("/follow_events", e.FollowEvents)
	eventGroup.GET("/created_by_me", e.FindEventCreatedByMe)
	eventGroup.GET("/tag", e.FindTagByToken)
	eventGroup.POST("/follow_tag", e.FollowTag)
	eventGroup.DELETE("/unfollow_tag", e.UnFollowTag)
	eventGroup.POST("/regis_event", e.RegisEvent)
	eventGroup.GET("/regis_events", e.RegisEvents)
	eventGroup.POST("/notifications", e.Notifications)
}

func NewEventController(db *pgxpool.Pool, firebaseService firebaseService.FireBaseService) *EventController {
	viper.AutomaticEnv()

	config := s3service.S3ServiceConfig{
		Region:             viper.GetString("AWS_REGION"),
		Bucket:             viper.GetString("AWS_BUCKET"),
		AwsAccessKeyID:     viper.GetString("AWS_ACCESS_KEY_ID"),
		AwsSecretAccessKey: viper.GetString("AWS_SECRET_ACCESS_KEY"),
	}

	newS3service, err := s3service.NewS3Service(&config)
	if err != nil {
		log.Fatal("Failed to initialize S3 service:", err)
	}

	return &EventController{
		es: eventservice.NewEventService(db),
		s3: newS3service,
		fb: firebaseService,
	}
}
