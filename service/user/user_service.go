package user

import (
	user "KUNoti/internal/controller/user/domain"
	"KUNoti/internal/controller/user/repository"
	"KUNoti/internal/request/userrequest"
	"KUNoti/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func (userService UserService) Create(ctx *gin.Context, createUserRequest userrequest.CreateUserRequest) (*user.User, error) {
	user, err := userService.userRepository.Create(ctx, createUserRequest)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (userService UserService) Login(ctx *gin.Context, loginUserRequest userrequest.LoginUserRequest) (*user.User, error) {
	user, err := userService.userRepository.FindUserByUsername(ctx, loginUserRequest.Username)
	if err != nil {
		return nil, err
	}

	if !CheckPasswordHash(loginUserRequest.Password, user.Password) {
		return nil, err
	}

	return &user, nil
}

func (userService UserService) Update(ctx *gin.Context, updateUserRequest userrequest.UpdateUserRequest) (*user.User, error) {
	user, err := userService.userRepository.Update(ctx, updateUserRequest)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (userService UserService) Delete(ctx *gin.Context, deleteUserRequest userrequest.DeleteUserRequest) (string, error) {
	id, err := userService.userRepository.Delete(ctx, deleteUserRequest)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (userService UserService) FindUserByID(ctx *gin.Context, findUserByIDRequest userrequest.FindUserByID) (*user.User, error) {
	user, err := userService.userRepository.FindUserByID(ctx, findUserByIDRequest)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func NewUserService(db *pgxpool.Pool) *UserService {
	queries := sqlc.New(db)
	return &UserService{
		userRepository: repository.NewUserRepository(db, queries),
	}
}
