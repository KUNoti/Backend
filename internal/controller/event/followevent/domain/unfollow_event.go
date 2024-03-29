package domain

type UnfollowEvent struct {
	Id      int `json:"id"`
	EventID int `json:"event_id"`
	UserID  int `json:"user_id"`
}
