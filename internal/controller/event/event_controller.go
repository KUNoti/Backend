package event

import (
	"KUNoti/internal/request/eventrequest"
	eventservice "KUNoti/service/event"
	"KUNoti/service/s3service"
	"github.com/spf13/viper"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EventController struct {
	es *eventservice.EventService
	s3 *s3service.S3Service
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

	imageURL, err := e.s3.Upload(s3service.EventImageFolder, createEventRequest.ImageFile)
	if err != nil {
		log.Println("Error saving image to S3:", err)
		ctx.JSON(http.StatusInternalServerError, "Error saving image")
		return
	}

	createEventRequest.Image = imageURL

	event, err := e.es.Create(ctx, createEventRequest)
	if err != nil {
		log.Println(err)
	}

	ctx.JSON(201, event)
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

func (e EventController) InitEndpoints(r *gin.RouterGroup) {
	eventGroup := r.Group("/event")
	eventGroup.POST("/create", e.CreateEvent)
	eventGroup.PUT("/update", e.UpdateEvent)
	eventGroup.DELETE("/delete", e.DeleteEvent)
	eventGroup.GET("", e.FinderEvent)
	eventGroup.GET("/events", e.Events)
}

func NewEventController(db *pgxpool.Pool) *EventController {
	viper.AutomaticEnv()

	config := s3service.S3ServiceConfig{
		Region:             viper.GetString("AWS_REGION"),
		Bucket:             viper.GetString("AWS_BUCKET"),
		AwsAccessKeyID:     viper.GetString("AWS_ACCESS_KEY_ID"),
		AwsSecretAccessKey: viper.GetString("AWS_SECRET_ACCESS_KEY"),
	}

	s3service, err := s3service.NewS3Service(&config)
	if err != nil {
		log.Fatal("Failed to initialize S3 service:", err)
	}

	return &EventController{
		es: eventservice.NewEventService(db),
		s3: s3service,
	}
}
