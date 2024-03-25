package userrequest

import (
	"KUNoti/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type UpdateUserRequest struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Name         string `json:"name"`
	SocialID     string `json:"social_id"`
	Email        string `json:"email"`
	ProfileImage string `json:"profile_image"`
	ID           int32  `json:"id"`
}

func CreateParamsFromUpdateUserRequest(cmd UpdateUserRequest) sqlc.UpdateUserByIDParams {
	params := sqlc.UpdateUserByIDParams{
		ID: cmd.ID,
		Name: pgtype.Text{
			String: cmd.Name,
			Valid:  cmd.Name != "",
		},
		SocialID: pgtype.Text{
			String: cmd.SocialID,
			Valid:  cmd.SocialID != "",
		},
		Email: pgtype.Text{
			String: cmd.Email,
			Valid:  cmd.Email != "",
		},
		ProfileImage: pgtype.Text{
			String: cmd.ProfileImage,
			Valid:  cmd.ProfileImage != "",
		},
	}
	return params
}
