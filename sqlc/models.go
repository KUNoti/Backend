// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package sqlc

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Event struct {
	ID           int32            `json:"id"`
	StartDate    pgtype.Timestamp `json:"start_date"`
	EndDate      pgtype.Timestamp `json:"end_date"`
	CreatedAt    pgtype.Timestamp `json:"created_at"`
	UpdatedAt    pgtype.Timestamp `json:"updated_at"`
	Title        string           `json:"title"`
	Latitude     float64          `json:"latitude"`
	Longitude    float64          `json:"longitude"`
	Price        float64          `json:"price"`
	Image        pgtype.Text      `json:"image"`
	Creator      string           `json:"creator"`
	Detail       string           `json:"detail"`
	LocationName string           `json:"location_name"`
	NeedRegis    bool             `json:"need_regis"`
	Tag          pgtype.Text      `json:"tag"`
}

type User struct {
	ID           int32            `json:"id"`
	Name         string           `json:"name"`
	Role         string           `json:"role"`
	CreatedAt    pgtype.Timestamp `json:"created_at"`
	UpdatedAt    pgtype.Timestamp `json:"updated_at"`
	Email        string           `json:"email"`
	ProfileImage pgtype.Text      `json:"profile_image"`
	Username     string           `json:"username"`
	Password     string           `json:"password"`
	SocialID     pgtype.Text      `json:"social_id"`
}
