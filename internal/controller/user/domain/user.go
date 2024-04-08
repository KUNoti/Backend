package domain

import "KUNoti/sqlc"

type User struct {
	ID           int32  `json:"id"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	Password     string
	Email        string `json:"email"`
	SocialID     string
	ProfileImage string `json:"profile_image"`
	Token        string `json:"token"`
}

func NewFromSqlc(u sqlc.User) User {
	user := User{
		ID:       u.ID,
		Name:     u.Name,
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
		Token:    u.Token,
	}
	if u.SocialID.Valid {
		user.SocialID = u.SocialID.String
	}
	if u.ProfileImage.Valid {
		user.ProfileImage = u.ProfileImage.String
	}
	return user
}
