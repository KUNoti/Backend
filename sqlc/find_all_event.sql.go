// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: find_all_event.sql

package sqlc

import (
	"context"
)

const findAllEvent = `-- name: FindAllEvent :many
SELECT id, start_date, end_date, created_at, updated_at, title, latitude, longitude, price, image, detail, location_name, need_regis, tag, creator, regis_amount, regis_max, need_noti
FROM events
WHERE regis_max > regis_amount
  AND start_date > CURRENT_TIMESTAMP AT TIME ZONE 'Asia/Bangkok'
ORDER BY id DESC
`

func (q *Queries) FindAllEvent(ctx context.Context) ([]Event, error) {
	rows, err := q.db.Query(ctx, findAllEvent)
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
			&i.RegisAmount,
			&i.RegisMax,
			&i.NeedNoti,
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
