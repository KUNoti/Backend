package domain

import (
	"KUNoti/sqlc"
	"time"
)

type FollowEvent struct {
	Id        int       `json:"id"`
	EventID   int       `json:"event_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewFromSqlc(fe sqlc.FollowingEvent) FollowEvent {
	event := FollowEvent{
		Id:        int(fe.ID),
		EventID:   int(fe.EventID),
		UserID:    int(fe.UserID),
		CreatedAt: fe.CreatedAt.Time,
		UpdatedAt: fe.UpdatedAt.Time,
	}
	return event
}
