// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: find_event_by_id.sql

package sqlc

import (
	"context"
)

const findEventByID = `-- name: FindEventByID :one
SELECT id, start_date, end_date, created_at, updated_at, title, latitude, longitude, price, image, detail, location_name, need_regis, tag, creator, regis_amount, regis_max, need_noti
FROM events
WHERE id = $1
ORDER BY id DESC
`

func (q *Queries) FindEventByID(ctx context.Context, id int32) (Event, error) {
	row := q.db.QueryRow(ctx, findEventByID, id)
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
