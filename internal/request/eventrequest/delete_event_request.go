package eventrequest

type DeleteEventRequest struct {
	ID int32 `json:"id"`
}

func CreateParamsFromDeleteRequest(cmd DeleteEventRequest) DeleteEventRequest {
	return DeleteEventRequest{
		ID: int32(cmd.ID),
	}
}
