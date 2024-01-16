package test

import (
	"KUNoti/internal/test/app"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetApplication(db *pgxpool.Pool) app.Application {
	return app.NewApplication(db)
}