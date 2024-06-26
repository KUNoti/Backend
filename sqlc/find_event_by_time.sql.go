// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: find_event_by_time.sql

package sqlc

import (
	"context"
)

const findEventByTime = `-- name: FindEventByTime :many
SELECT id, start_date, end_date, created_at, updated_at, title, latitude, longitude, price, image, detail, location_name, need_regis, tag, creator, regis_amount, regis_max, need_noti
FROM events
WHERE start_date BETWEEN
        (CURRENT_TIMESTAMP AT TIME ZONE 'Asia/Bangkok') - interval '5 minutes'
  AND
    (CURRENT_TIMESTAMP AT TIME ZONE 'Asia/Bangkok') + interval '5 minutes'
  AND
    need_noti = true
ORDER BY id DESC
`

func (q *Queries) FindEventByTime(ctx context.Context) ([]Event, error) {
	rows, err := q.db.Query(ctx, findEventByTime)
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
