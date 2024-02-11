package domain

import (
	"KUNoti/sqlc"
	"time"
)

type Event struct {
	Id           int
	Title        string
	Lat          float64
	Lon          float64
	StartDate    time.Time
	EndDate      time.Time
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Price        float64   `json:"price"`
	Rating       float64   `json:"rating"`
	Image        string    `json:"image"`
	Creator      string    `json:"creator"`
	Detail       string    `json:"detail"`
	LocationName string    `json:"location_name"`
	NeedRegis    bool      `json:"need_regis"`
}

func NewFromSqlc(e sqlc.Event) Event {
	return Event{
		Id:           int(e.ID),
		Title:        e.Title,
		Lat:          e.Latitude,
		Lon:          e.Longitude,
		StartDate:    e.StartDate.Time,
		EndDate:      e.EndDate.Time,
		CreatedAt:    e.CreatedAt.Time,
		UpdatedAt:    e.UpdatedAt.Time,
		Price:        e.Price,
		Rating:       e.Rating,
		Image:        e.Image,
		Creator:      e.Creator,
		Detail:       e.Detail,
		LocationName: e.LocationName,
		NeedRegis:    e.NeedRegis,
	}
}
