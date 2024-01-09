package user

import (
	"KUNoti/internal/user/app"
	"github.com/jackc/pgx/v5/pgxpool"
)

// type Server struct {
// 	router *gin.Engine
// }

func GetApplication(db *pgxpool.Pool) app.Application {
	return app.NewApplication(db)
}
