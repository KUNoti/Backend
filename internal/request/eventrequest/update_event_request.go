package eventrequest

import (
	"KUNoti/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type UpdateEventRequest struct {
	Title        string    `json:"title"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	Price        float64   `json:"price"`
	Creator      int       `json:"creator"`
	Detail       string    `json:"detail"`
	LocationName string    `json:"location_name"`
	NeedRegis    *bool     `json:"need_regis"`
	ID           int32     `json:"id"`
	Tag          string    `json:"tag"`
	RegisAmount  int       `json:"regis_amount"`
	RegisMax     int       `json:"regis_max"`
}

func CreateParamsFromUpdateRequest(cmd UpdateEventRequest) sqlc.UpdateEventByIDParams {
	params := sqlc.UpdateEventByIDParams{
		ID: cmd.ID,
		Title: pgtype.Text{
			String: cmd.Title,
			Valid:  cmd.Title != "",
		},
		Latitude: pgtype.Float8{
			Float64: cmd.Latitude,
			Valid:   cmd.Latitude != 0.0,
		},
		Longitude: pgtype.Float8{
			Float64: cmd.Longitude,
			Valid:   cmd.Longitude != 0.0,
		},
		StartDate: pgtype.Timestamp{
			Time:  cmd.StartDate,
			Valid: cmd.StartDate != time.Time{},
		},
		EndDate: pgtype.Timestamp{
			Time:  cmd.EndDate,
			Valid: cmd.EndDate != time.Time{},
		},
		Price: pgtype.Float8{
			Float64: cmd.Price,
			Valid:   cmd.Price != 0.0,
		},
		Creator: pgtype.Int4{
			Int32: int32(cmd.Creator),
			Valid: cmd.Creator != 0,
		},
		Detail: pgtype.Text{
			String: cmd.Detail,
			Valid:  cmd.Detail != "",
		},
		LocationName: pgtype.Text{
			String: cmd.LocationName,
			Valid:  cmd.LocationName != "",
		},
		Tag: pgtype.Text{
			String: cmd.Tag,
			Valid:  cmd.Tag != "",
		},
		RegisAmount: pgtype.Int4{
			Int32: int32(cmd.RegisAmount),
			Valid: cmd.RegisAmount != 0,
		},
		RegisMax: pgtype.Int4{
			Int32: int32(cmd.RegisMax),
			Valid: cmd.RegisMax != 0,
		},
	}
	if cmd.NeedRegis != nil {
		params.NeedRegis.Bool = *cmd.NeedRegis
		params.NeedRegis.Valid = true
	}
	return params
}
