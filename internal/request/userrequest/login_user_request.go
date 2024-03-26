package userrequest

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
