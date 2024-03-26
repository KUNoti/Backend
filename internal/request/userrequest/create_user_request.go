package userrequest

import (
	"KUNoti/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
	"mime/multipart"
	"time"
)

type CreateUserRequest struct {
	Username     string                `form:"username"`
	Password     string                `form:"password"`
	Name         string                `form:"name"`
	SocialID     string                `form:"social_id"`
	Email        string                `form:"email"`
	ProfileImage string                `form:"profile_image"`
	ProfileFile  *multipart.FileHeader `form:"profile_file"`
}

func CreateParamsFromCreateUserRequest(cmd CreateUserRequest) sqlc.CreateUserParams {
	return sqlc.CreateUserParams{
		Name:     cmd.Name,
		Username: cmd.Username,
		Password: cmd.Password,
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
	Email        string
	ProfileImage string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
