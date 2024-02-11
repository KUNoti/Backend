package event

import (
	"KUNoti/internal/request/eventrequest"
	eventservice "KUNoti/service/event"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EventController struct {
	es *eventservice.EventService
}

func (e EventController) CreateEvent(ctx *gin.Context) {
	// Get body from req

	// service.CreateEvent -> service.repo.CreateEvent
	var createEventRequest eventrequest.CreateEventRequest
	err := ctx.BindJSON(&createEventRequest)
	log.Println("test: ", err)
	if err != nil {
		log.Println(err.Error())
		log.Printf("Error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	log.Println(createEventRequest)

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

func (e EventController) HiEvent(ctx *gin.Context) {
	ctx.JSON(200, "hi event")
}

func (e EventController) InitEndpoints(r *gin.RouterGroup) {
	eventGroup := r.Group("/event")
	eventGroup.GET("/hi", e.HiEvent)
	eventGroup.POST("/create", e.CreateEvent)
	eventGroup.PUT("/update", e.UpdateEvent)
	eventGroup.DELETE("/delete", e.DeleteEvent)
}

func NewEventController(db *pgxpool.Pool) *EventController {
	return &EventController{
		es: eventservice.NewEventService(db),
	}
}
