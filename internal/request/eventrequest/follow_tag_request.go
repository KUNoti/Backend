package eventrequest

import "KUNoti/sqlc"

type FollowTagRequest struct {
	Tag       string `json:"tag"`
	UserToken string `json:"user_token"`
}

func CreateParamsFromFollowTagRequest(cmd FollowTagRequest) sqlc.FollowTagParams {
	return sqlc.FollowTagParams{
		Tag:       cmd.Tag,
		UserToken: cmd.UserToken,
	}
}
