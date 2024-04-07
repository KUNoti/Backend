package repository

import (
	event "KUNoti/internal/controller/event/domain"
	followByTag "KUNoti/internal/controller/event/followbytag/domain"
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

func (er EventRepository) FollowEvents(ctx *gin.Context, userID int) ([]event.Event, error) {
	followEvents, err := er.queries.FindAllFollowEvent(ctx, int32(userID))
	if err != nil {
		return nil, err
	}
	events := make([]event.Event, len(followEvents))

	for i, e := range followEvents {
		eventSqlc, err := er.queries.FindEventByID(ctx, e.EventID)
		if err != nil {
			return nil, err
		}
		events[i] = event.NewFromSqlc(eventSqlc)
	}
	return events, nil
}

func (er EventRepository) FindEventCreatedByID(ctx *gin.Context, userID int) ([]event.Event, error) {
	createdEventByMe, err := er.queries.FindEventCreatedByID(ctx, int32(userID))
	if err != nil {
		return nil, err
	}
	events := make([]event.Event, len(createdEventByMe))
	for i, e := range createdEventByMe {
		events[i] = event.NewFromSqlc(e)
	}
	return events, nil
}

func (er EventRepository) FollowTag(ctx *gin.Context, request eventrequest.FollowTagRequest) (followByTag.FollowByTag, error) {
	arg := eventrequest.CreateParamsFromFollowTagRequest(request)
	followTagSqlc, err := er.queries.FollowTag(ctx, arg)
	if err != nil {
		return followByTag.FollowByTag{}, err
	}
	followConvert := followByTag.NewFromSqlc(followTagSqlc)
	return followConvert, nil
}

func (er EventRepository) UnfollowTag(ctx *gin.Context, request eventrequest.UnFollowTagRequest) (string, error) {
	arg := eventrequest.CreateParamsFromUnFollowTagRequest(request)

	err := er.queries.UnfollowTag(ctx, arg)
	if err != nil {
		return "", err
	}
	return "unfollow success", nil
}

func (er EventRepository) FindTokensByTagName(ctx *gin.Context, tag string) ([]string, error) {
	tokens, err := er.queries.FindTokensByTagName(ctx, tag)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

func (er EventRepository) RegisEvent(ctx *gin.Context, request eventrequest.RegisEventRequest) (string, error) {
	arg := eventrequest.CrateParamsFromRegisRequest(request)
	_, err := er.queries.RegisEventByID(ctx, arg.EventID)
	if err != nil {
		return "", err
	}
	_, err = er.queries.CreateRegisEvent(ctx, arg)
	if err != nil {
		return "", err
	}
	return "regis success", nil
}

func (er EventRepository) FindRegisEventByUserID(ctx *gin.Context, userID int) ([]event.Event, error) {
	regisEventByMe, err := er.queries.FindRegisEventByUserID(ctx, int32(userID))
	if err != nil {
		return nil, err
	}
	events := make([]event.Event, len(regisEventByMe))
	for i, re := range regisEventByMe {
		e, err := er.queries.FindEventByID(ctx, re.EventID)
		if err != nil {
			return nil, err
		}
		events[i] = event.NewFromSqlc(e)
	}
	return events, nil
}

func NewEventRepository(db *pgxpool.Pool, queries *sqlc.Queries) *EventRepository {
	return &EventRepository{
		DB:      db,
		queries: queries,
	}
}
