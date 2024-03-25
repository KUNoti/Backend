package repository

import (
	user "KUNoti/internal/controller/user/domain"
	"KUNoti/internal/request/userrequest"
	"KUNoti/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

type UserRepository struct {
	DB      *pgxpool.Pool
	queries *sqlc.Queries
}

func (ur UserRepository) Create(ctx *gin.Context, createUserRequest userrequest.CreateUserRequest) (user.User, error) {
	hashedPassword, err := hashPassword(createUserRequest.Password)
	if err != nil {
		return user.User{}, err
	}
	if hashedPassword != "" {
		createUserRequest.Password = hashedPassword
	}
	arg := userrequest.CreateParamsFromCreateUserRequest(createUserRequest)

	userSqlc, err := ur.queries.CreateUser(ctx, arg)
	if err != nil {
		return user.User{}, err
	}
	userConvert := user.NewFromSqlc(userSqlc)
	return userConvert, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (ur UserRepository) Update(ctx *gin.Context, updateUserRequest userrequest.UpdateUserRequest) (user.User, error) {
	arg := userrequest.CreateParamsFromUpdateUserRequest(updateUserRequest)

	userSqlc, err := ur.queries.UpdateUserByID(ctx, arg)
	if err != nil {
		return user.User{}, err
	}
	userConvert := user.NewFromSqlc(userSqlc)
	return userConvert, nil
}

func (ur UserRepository) Delete(ctx *gin.Context, deleteUserRequest userrequest.DeleteUserRequest) (string, error) {
	arg := userrequest.CreateParamsFromDeleteUserRequest(deleteUserRequest)

	id, err := ur.queries.DeleteUser(ctx, arg.ID)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(id)), nil
}

func (ur UserRepository) FindUserByID(ctx *gin.Context, finderUserRequest userrequest.FindUserByID) (user.User, error) {
	userSqlc, err := ur.queries.FindUserByID(ctx, finderUserRequest.ID)
	if err != nil {
		return user.User{}, err
	}
	userConvert := user.NewFromSqlc(userSqlc)
	return userConvert, nil
}

func NewUserRepository(db *pgxpool.Pool, queries *sqlc.Queries) *UserRepository {
	return &UserRepository{
		DB:      db,
		queries: queries,
	}
}
