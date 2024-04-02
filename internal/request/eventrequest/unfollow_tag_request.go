package eventrequest

import "KUNoti/sqlc"

type UnFollowTagRequest struct {
	Tag       string `json:"tag"`
	UserToken string `json:"user_token"`
}

func CreateParamsFromUnFollowTagRequest(cmd UnFollowTagRequest) sqlc.UnfollowTagParams {
	return sqlc.UnfollowTagParams{
		Tag:       cmd.Tag,
		UserToken: cmd.UserToken,
	}
}
