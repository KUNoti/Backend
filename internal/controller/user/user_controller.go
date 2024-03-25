package user

import (
	"KUNoti/internal/request/userrequest"
	userservice "KUNoti/service/user"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"
	"strconv"
)

type UserController struct {
	us *userservice.UserService
}

func (u UserController) CreateUser(ctx *gin.Context) {
	var createUserRequest userrequest.CreateUserRequest
	err := ctx.BindJSON(&createUserRequest)
	if err != nil {
		log.Println(err.Error())
		log.Printf("Error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := u.us.Create(ctx, createUserRequest)
	if err != nil {
		log.Println(err.Error())
		log.Printf("Error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(201, user)
}

func (u UserController) UpdateUser(ctx *gin.Context) {
	var updateUserRequest userrequest.UpdateUserRequest
	err := ctx.BindJSON(&updateUserRequest)
	if err != nil {
		log.Println(err.Error())
		log.Printf("Error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := u.us.Update(ctx, updateUserRequest)
	if err != nil {
		log.Println(err.Error())
		log.Printf("Error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(200, user)
}

func (u UserController) DeleteUser(ctx *gin.Context) {
	var deleteUserRequest userrequest.DeleteUserRequest
	err := ctx.BindJSON(&deleteUserRequest)
	if err != nil {
		log.Println(err.Error())
		log.Printf("Error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := u.us.Delete(ctx, deleteUserRequest)
	if err != nil {
		log.Println(err.Error())
		log.Printf("Error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(200, "delete event ID : "+id)
}

func (u UserController) User(ctx *gin.Context) {
	var findUserByIDRequest userrequest.FindUserByID
	if queryParam, ok := ctx.GetQuery("id"); ok {
		id, err := strconv.Atoi(queryParam)
		if err != nil {
			log.Println(err.Error())
			log.Printf("Error: %v\n", err)
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		findUserByIDRequest.ID = int32(id)
	}
	user, err := u.us.FindUserByID(ctx, findUserByIDRequest)
	if err != nil {
		log.Println(err.Error())
		log.Printf("Error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(200, user)
}

func (u UserController) InitEndpoints(r *gin.RouterGroup) {
	userGroup := r.Group("/user")
	userGroup.GET("", u.User)
	userGroup.POST("/create", u.CreateUser)
	userGroup.PUT("/update", u.UpdateUser)
	userGroup.DELETE("/delete", u.DeleteUser)
}

func NewUserController(db *pgxpool.Pool) *UserController {
	return &UserController{
		us: userservice.NewUserService(db),
	}
}
