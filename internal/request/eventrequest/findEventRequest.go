package eventrequest

import (
	"KUNoti/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type FinderEventRequest struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Price        string    `json:"price"`
	Creator      string    `json:"creator"`
	LocationName string    `json:"location_name"`
	NeedRegis    string    `json:"need_regis"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
}

func CreateParamsFromFinderRequest(cmd FinderEventRequest) sqlc.FinderEventParams {
	return sqlc.FinderEventParams{
		Title: pgtype.Text{
			String: cmd.Title,
			Valid:  cmd.Title != "",
		},
		Creator: pgtype.Text{
			String: cmd.Creator,
			Valid:  cmd.Creator != "",
		},
		LocationName: pgtype.Text{
			String: cmd.LocationName,
			Valid:  cmd.LocationName != "",
		},
	 }
}
