package eventrequest

import "KUNoti/sqlc"

type UnfollowEventRequest struct {
	UserID  int `json:"user_id"`
	EventID int `json:"event_id"`
}

func CrateParamsFromUnfollowRequest(cmd UnfollowEventRequest) sqlc.UnfollowEventParams {
	return sqlc.UnfollowEventParams{
		EventID: int32(cmd.EventID),
		UserID:  int32(cmd.UserID),
	}
}
