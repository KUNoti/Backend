package eventrequest

import (
	"KUNoti/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type CreateEventRequest struct {
	Title        string    `json:"title"`
	Lat          float64   `json:"latitude"`
	Lon          float64   `json:"longitude"`
	StartDate    time.Time `json:"start_date_time"`
	EndDate      time.Time `json:"end_date_time"`
	Price        float64   `json:"price"`
	Image        string    `json:"image"`
	Creator      string    `json:"creator"`
	Detail       string    `json:"detail"`
	LocationName string    `json:"location_name"`
	NeedRegis    bool      `json:"need_regis"`
}

func CreateParamsFromCreateRequest(cmd CreateEventRequest) sqlc.CreateEventParams {
	return sqlc.CreateEventParams{
		Title:     cmd.Title,
		Latitude:  cmd.Lat,
		Longitude: cmd.Lon,
		StartDate: pgtype.Timestamp{
			Time:  cmd.StartDate,
			Valid: true,
		},
		EndDate: pgtype.Timestamp{
			Time:  cmd.EndDate,
			Valid: true,
		},
		Price: cmd.Price,
		Image: pgtype.Text{
			String: cmd.Image,
			Valid:  cmd.Image != "",
		},
		Creator:      cmd.Creator,
		Detail:       cmd.Detail,
		LocationName: cmd.LocationName,
		NeedRegis:    cmd.NeedRegis,
	}
}
