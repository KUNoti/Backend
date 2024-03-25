package userrequest

import (
	"KUNoti/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type CreateUserRequest struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Name         string `json:"name"`
	SocialID     string `json:"social_id"`
	Role         string `json:"role"`
	Email        string `json:"email"`
	ProfileImage string `json:"profile_image"`
}

func CreateParamsFromCreateUserRequest(cmd CreateUserRequest) sqlc.CreateUserParams {
	return sqlc.CreateUserParams{
		Name:     cmd.Name,
		Username: cmd.Username,
		Password: cmd.Password,
		Role:     "User",
		Email:    cmd.Email,
		SocialID: pgtype.Text{
			String: cmd.SocialID,
			Valid:  cmd.SocialID != "",
		},
		ProfileImage: pgtype.Text{
			String: cmd.ProfileImage,
			Valid:  cmd.ProfileImage != "",
		},
	}
}

type User struct {
	ID           int
	Username     string
	Password     string
	Name         string
	SocialID     string
	Role         string
	Email        string
	ProfileImage string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
