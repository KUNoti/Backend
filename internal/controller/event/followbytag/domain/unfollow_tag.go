package domain

type UnFollowTag struct {
	Id        int    `json:"id"`
	Tag       string `json:"tag"`
	UserToken string `json:"user_token"`
}
