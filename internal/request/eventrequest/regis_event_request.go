package eventrequest

import "KUNoti/sqlc"

type RegisEventRequest struct {
	UserID  int `json:"user_id"`
	EventID int `json:"event_id"`
}

func CrateParamsFromRegisRequest(cmd RegisEventRequest) sqlc.CreateRegisEventParams {
	return sqlc.CreateRegisEventParams{
		EventID: int32(cmd.EventID),
		UserID:  int32(cmd.UserID),
	}
}
