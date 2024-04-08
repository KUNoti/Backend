package eventservice

import (
	event "KUNoti/internal/controller/event/domain"
	followtag "KUNoti/internal/controller/event/followbytag/domain"
	followevent "KUNoti/internal/controller/event/followevent/domain"
	"KUNoti/internal/controller/event/repository"
	"KUNoti/internal/request/eventrequest"
	"KUNoti/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
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

func (eventService EventService) FindFollowEvent(ctx *gin.Context, userID eventrequest.FindFollowEventRequest) ([]event.Event, error) {
	events, err := eventService.eventRepository.FollowEvents(ctx, userID.UserID)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (eventService EventService) FindEventCreatedByMe(ctx *gin.Context, userID eventrequest.FindEventCreatedByMeRequest) ([]event.Event, error) {
	events, err := eventService.eventRepository.FindEventCreatedByID(ctx, userID.UserID)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (eventService EventService) FollowTag(ctx *gin.Context, request eventrequest.FollowTagRequest) (followtag.FollowByTag, error) {
	followT, err := eventService.eventRepository.FollowTag(ctx, request)
	if err != nil {
		return followtag.FollowByTag{}, err
	}
	return followT, nil
}

func (eventService EventService) UnfollowTag(ctx *gin.Context, request eventrequest.UnFollowTagRequest) (string, error) {
	tag, err := eventService.eventRepository.UnfollowTag(ctx, request)
	if err != nil {
		return "", err
	}
	return tag, nil
}

func (eventService EventService) FindTokensByTagName(ctx *gin.Context, tag string) ([]string, error) {
	tokens, err := eventService.eventRepository.FindTokensByTagName(ctx, tag)
	if err != nil {
		log.Printf("Error finding tokens by tag: %v\n", err)
		return nil, err // Return nil slice and the error
	}
	return tokens, nil // Return the fetched tokens and nil as error
}

func (eventService EventService) FindTagByToken(ctx *gin.Context, request eventrequest.FindTagByToken) ([]string, error) {
	tag, err := eventService.eventRepository.FindTagByToken(ctx, request.Token)
	if err != nil {
		log.Printf("Error finding tokens by tag: %v\n", err)
		return nil, err
	}
	return tag, err
}

func NewEventService(db *pgxpool.Pool) *EventService {
	queries := sqlc.New(db)
	return &EventService{
		eventRepository: repository.NewEventRepository(db, queries),
	}
}
