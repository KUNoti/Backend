package eventrequest

import (
	"KUNoti/sqlc"
	"mime/multipart"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type CreateEventRequest struct {
	Title        string                `form:"title"`
	Lat          float64               `form:"latitude"`
	Lon          float64               `form:"longitude"`
	StartDate    time.Time             `form:"start_date_time"`
	EndDate      time.Time             `form:"end_date_time"`
	Price        float64               `form:"price"`
	Image        string                `form:"image"`
	Creator      int                   `form:"creator"`
	Detail       string                `form:"detail"`
	LocationName string                `form:"location_name"`
	NeedRegis    bool                  `form:"need_regis"`
	ImageFile    *multipart.FileHeader `form:"image_file"`
	//Tag [list tag]
	//User id
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
		Creator:      int32(cmd.Creator),
		Detail:       cmd.Detail,
		LocationName: cmd.LocationName,
		NeedRegis:    cmd.NeedRegis,
	}
}
