package eventrequest

import "KUNoti/sqlc"

type FollowEventRequest struct {
	UserID  int `json:"user_id"`
	EventID int `json:"event_id"`
}

func CrateParamsFromFollowRequest(cmd FollowEventRequest) sqlc.FollowEventParams {
	return sqlc.FollowEventParams{
		EventID: int32(cmd.EventID),
		UserID:  int32(cmd.UserID),
	}
}
