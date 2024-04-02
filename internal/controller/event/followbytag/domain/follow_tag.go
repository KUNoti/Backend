package domain

import (
	"KUNoti/sqlc"
	"time"
)

type FollowByTag struct {
	Id        int       `json:"id"`
	Tag       string    `json:"tag"`
	UserToken string    `json:"user_token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewFromSqlc(ft sqlc.FollowByTag) FollowByTag {
	tag := FollowByTag{
		Id:        int(ft.ID),
		Tag:       ft.Tag,
		UserToken: ft.UserToken,
		CreatedAt: ft.CreatedAt.Time,
		UpdatedAt: ft.UpdatedAt.Time,
	}
	return tag
}
