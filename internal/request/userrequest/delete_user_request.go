package userrequest

type DeleteUserRequest struct {
	ID int32 `json:"id"`
}

func CreateParamsFromDeleteUserRequest(cmd DeleteUserRequest) DeleteUserRequest {
	return DeleteUserRequest{
		ID: int32(cmd.ID),
	}
}
