package eventservice

import (
	event "KUNoti/internal/controller/event/domain"
	followevent "KUNoti/internal/controller/event/followevent/domain"
	"KUNoti/internal/controller/event/repository"
	"KUNoti/internal/request/eventrequest"
	"KUNoti/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EventService struct {
	eventRepository *repository.EventRepository
}

func (eventService EventService) Create(ctx *gin.Context, createEventRequest eventrequest.CreateEventRequest) (*event.Event, error) {
	event, err := eventService.eventRepository.Create(ctx, createEventRequest)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (eventService EventService) Update(ctx *gin.Context, updateEventRequest eventrequest.UpdateEventRequest) (*event.Event, error) {
	event, err := eventService.eventRepository.Update(ctx, updateEventRequest)
	if err != nil {
		return nil, err
	}
	return &event, err
}

func (eventService EventService) Delete(ctx *gin.Context, deleteEventRequest eventrequest.DeleteEventRequest) (string, error) {
	id, err := eventService.eventRepository.Delete(ctx, deleteEventRequest)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (eventService EventService) Finder(ctx *gin.Context, finderEventRequest eventrequest.FinderEventRequest) ([]event.Event, error) {
	events, err := eventService.eventRepository.Finder(ctx, finderEventRequest)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (eventService EventService) FindAll(ctx *gin.Context) ([]event.Event, error) {
	events, err := eventService.eventRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (eventService EventService) Follow(ctx *gin.Context, followEventRequest eventrequest.FollowEventRequest) (followevent.FollowEvent, error) {
	followE, err := eventService.eventRepository.FollowEvent(ctx, followEventRequest)
	if err != nil {
		return followevent.FollowEvent{}, err
	}
	return followE, nil
}

func (eventService EventService) Unfollow(ctx *gin.Context, unfollowEventRequest eventrequest.UnfollowEventRequest) (string, error) {
	eventID, err := eventService.eventRepository.UnfollowEvent(ctx, unfollowEventRequest)
	if err != nil {
		return "", err
	}
	return eventID, nil
}

func NewEventService(db *pgxpool.Pool) *EventService {
	queries := sqlc.New(db)
	return &EventService{
		eventRepository: repository.NewEventRepository(db, queries),
	}
}
