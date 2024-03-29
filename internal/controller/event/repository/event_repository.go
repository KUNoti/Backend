package repository

import (
	event "KUNoti/internal/controller/event/domain"
	followevent "KUNoti/internal/controller/event/followevent/domain"
	"KUNoti/internal/request/eventrequest"
	"KUNoti/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"strconv"
)

type EventRepository struct {
	DB      *pgxpool.Pool
	queries *sqlc.Queries
}

func (er EventRepository) Create(ctx *gin.Context, createEventRequest eventrequest.CreateEventRequest) (event.Event, error) {
	arg := eventrequest.CreateParamsFromCreateRequest(createEventRequest)

	eventSqlc, err := er.queries.CreateEvent(ctx, arg)
	if err != nil {
		return event.Event{}, err
	}
	eventConvert := event.NewFromSqlc(eventSqlc)
	return eventConvert, nil
}

func (er EventRepository) Update(ctx *gin.Context, updateEventRequest eventrequest.UpdateEventRequest) (event.Event, error) {
	arg := eventrequest.CreateParamsFromUpdateRequest(updateEventRequest)

	eventSqlc, err := er.queries.UpdateEventByID(ctx, arg)
	if err != nil {
		return event.Event{}, err
	}
	eventConvert := event.NewFromSqlc(eventSqlc)
	return eventConvert, nil
}

func (er EventRepository) Delete(ctx *gin.Context, deleteEventRequest eventrequest.DeleteEventRequest) (string, error) {
	arg := eventrequest.CreateParamsFromDeleteRequest(deleteEventRequest)

	id, err := er.queries.DeleteEventByID(ctx, arg.ID)
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func (er EventRepository) Finder(ctx *gin.Context, finderEventRequest eventrequest.FinderEventRequest) ([]event.Event, error) {
	eventSqlcs, err := er.queries.FinderEvent(ctx, finderEventRequest.Keyword)
	if err != nil {
		return nil, err
	}
	eventConvert := make([]event.Event, len(eventSqlcs))
	for i, e := range eventSqlcs {
		eventConvert[i] = event.NewFromSqlc(e)
	}
	return eventConvert, nil
}

func (er EventRepository) FindAll(ctx *gin.Context) ([]event.Event, error) {
	eventSqlcs, err := er.queries.FindAllEvent(ctx)
	if err != nil {
		return nil, err
	}
	eventConvert := make([]event.Event, len(eventSqlcs))
	for i, e := range eventSqlcs {
		eventConvert[i] = event.NewFromSqlc(e)
	}
	return eventConvert, nil
}

func (er EventRepository) FollowEvent(ctx *gin.Context, followEventRequest eventrequest.FollowEventRequest) (followevent.FollowEvent, error) {
	arg := eventrequest.CrateParamsFromFollowRequest(followEventRequest)

	followEventSqlcs, err := er.queries.FollowEvent(ctx, arg)
	if err != nil {
		return followevent.FollowEvent{}, err
	}
	followConvert := followevent.NewFromSqlc(followEventSqlcs)
	return followConvert, nil
}

func (er EventRepository) UnfollowEvent(ctx *gin.Context, unfollow eventrequest.UnfollowEventRequest) (string, error) {
	arg := eventrequest.CrateParamsFromUnfollowRequest(unfollow)

	eventId, err := er.queries.UnfollowEvent(ctx, arg)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(eventId)), nil
}

func NewEventRepository(db *pgxpool.Pool, queries *sqlc.Queries) *EventRepository {
	return &EventRepository{
		DB:      db,
		queries: queries,
	}
}
