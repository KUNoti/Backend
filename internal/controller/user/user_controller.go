package user

import (
	"KUNoti/internal/request/userrequest"
	"KUNoti/service/s3service"
	userservice "KUNoti/service/user"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strconv"
)

type UserController struct {
	us *userservice.UserService
	s3 *s3service.S3Service
}

func (u UserController) CreateUser(ctx *gin.Context) {
	var createUserRequest userrequest.CreateUserRequest
	err := ctx.ShouldBind(&createUserRequest)
	if err != nil {
		log.Println(err.Error())
		log.Printf("Error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	imageURL, err := u.s3.Upload(s3service.UserImageFolder, createUserRequest.ProfileFile)
	if err != nil {
		log.Println("Error saving image to S3:", err)
		ctx.JSON(http.StatusInternalServerError, "Error saving image")
		return
	}

	createUserRequest.ProfileImage = imageURL

	if err != nil {
		log.Println("Error saving image to S3:", err)
		ctx.JSON(http.StatusInternalServerError, "Error saving image")
		return
	}

	_, err = u.us.Create(ctx, createUserRequest)
	if err != nil {
		log.Println(err.Error())
		log.Printf("Error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(201, nil)
}

func (u UserController) Login(ctx *gin.Context) {
	var loginUserRequest userrequest.LoginUserRequest
	err := ctx.ShouldBind(&loginUserRequest)

	user, err := u.us.Login(ctx, loginUserRequest)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Password = ""
	ctx.JSON(200, user)
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
	userGroup.POST("/login", u.Login)
}

func NewUserController(db *pgxpool.Pool) *UserController {
	viper.AutomaticEnv()

	config := s3service.S3ServiceConfig{
		Region:             viper.GetString("AWS_REGION"),
		Bucket:             viper.GetString("AWS_BUCKET"),
		AwsAccessKeyID:     viper.GetString("AWS_ACCESS_KEY_ID"),
		AwsSecretAccessKey: viper.GetString("AWS_SECRET_ACCESS_KEY"),
	}

	s3service, err := s3service.NewS3Service(&config)
	if err != nil {
		log.Fatal("Failed to initialize S3 service:", err)
	}

	return &UserController{
		us: userservice.NewUserService(db),
		s3: s3service,
	}
}
