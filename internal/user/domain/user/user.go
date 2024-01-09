package user

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// type Server struct {
// 	router *gin.Engine
// }

type Test struct {
	ID int
}

type UserResponse struct {
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

type LoginUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        UserResponse `json:"user"`
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func LoginUser(ctx *gin.Context) {
	var req LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// todo

	// mock
	ctx.JSON(http.StatusOK, gin.H{"test": "test"})
}
