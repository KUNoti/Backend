package domain

import (
	"KUNoti/sqlc"
	"time"
)

type Event struct {
	Id           int       `json:"id"`
	Title        string    `json:"title"`
	Lat          float64   `json:"latitude"`
	Lon          float64   `json:"longitude"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Price        float64   `json:"price"`
	Rating       float64   `json:"rating"`
	Image        string    `json:"image"`
	Creator      int       `json:"creator"`
	Detail       string    `json:"detail"`
	LocationName string    `json:"location_name"`
	NeedRegis    bool      `json:"need_regis"`
}

func NewFromSqlc(e sqlc.Event) Event {
	event := Event{
		Id:           int(e.ID),
		Title:        e.Title,
		Lat:          e.Latitude,
		Lon:          e.Longitude,
		StartDate:    e.StartDate.Time,
		EndDate:      e.EndDate.Time,
		CreatedAt:    e.CreatedAt.Time,
		UpdatedAt:    e.UpdatedAt.Time,
		Price:        e.Price,
		Creator:      int(e.Creator),
		Detail:       e.Detail,
		LocationName: e.LocationName,
		NeedRegis:    e.NeedRegis,
	}
	if e.Image.Valid {
		event.Image = e.Image.String
	}
	return event
}
