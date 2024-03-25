package domain

import "KUNoti/sqlc"

type User struct {
	Name         string
	Username     string
	Password     string
	Email        string
	SocialID     string
	ProfileImage string
}

func NewFromSqlc(u sqlc.User) User {
	user := User{
		Name:     u.Name,
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
	}
	if u.SocialID.Valid {
		user.SocialID = u.SocialID.String
	}
	if u.ProfileImage.Valid {
		user.ProfileImage = u.ProfileImage.String
	}
	return user
}
