// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: regis_event.sql

package sqlc

import (
	"context"
)

const createRegisEvent = `-- name: CreateRegisEvent :one
INSERT INTO "regis_events"
(
    event_id, user_id, created_at, updated_at
)
VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
    RETURNING id, event_id, user_id, created_at, updated_at
`

type CreateRegisEventParams struct {
	EventID int32 `json:"event_id"`
	UserID  int32 `json:"user_id"`
}

func (q *Queries) CreateRegisEvent(ctx context.Context, arg CreateRegisEventParams) (RegisEvent, error) {
	row := q.db.QueryRow(ctx, createRegisEvent, arg.EventID, arg.UserID)
	var i RegisEvent
	err := row.Scan(
		&i.ID,
		&i.EventID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const regisEventByID = `-- name: RegisEventByID :one
UPDATE events
SET regis_amount = CASE
                       WHEN regis_max > regis_amount THEN regis_amount + 1
                       ELSE regis_amount
    END,
    updated_at = CASE
                     WHEN regis_max > regis_amount THEN CURRENT_TIMESTAMP
                     ELSE updated_at
    END
WHERE id = $1
      RETURNING id, start_date, end_date, created_at, updated_at, title, latitude, longitude, price, image, detail, location_name, need_regis, tag, creator, regis_amount, regis_max, need_noti
`

func (q *Queries) RegisEventByID(ctx context.Context, id int32) (Event, error) {
	row := q.db.QueryRow(ctx, regisEventByID, id)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.StartDate,
		&i.EndDate,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.Latitude,
		&i.Longitude,
		&i.Price,
		&i.Image,
		&i.Detail,
		&i.LocationName,
		&i.NeedRegis,
		&i.Tag,
		&i.Creator,
		&i.RegisAmount,
		&i.RegisMax,
		&i.NeedNoti,
	)
	return i, err
}
