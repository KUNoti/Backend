// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: find_event_created_by_id.sql

package sqlc

import (
	"context"
)

const findEventCreatedByID = `-- name: FindEventCreatedByID :many
SELECT id, start_date, end_date, created_at, updated_at, title, latitude, longitude, price, image, detail, location_name, need_regis, tag, creator
FROM events
WHERE creator = $1
ORDER BY id DESC
`

func (q *Queries) FindEventCreatedByID(ctx context.Context, creator int32) ([]Event, error) {
	rows, err := q.db.Query(ctx, findEventCreatedByID, creator)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Event{}
	for rows.Next() {
		var i Event
		if err := rows.Scan(
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
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
