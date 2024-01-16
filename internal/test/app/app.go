package app

import "github.com/jackc/pgx/v5/pgxpool"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
}

type Queries struct {
}

func NewApplication(db *pgxpool.Pool) Application {
	// queries := sqlc.New(db)

	return Application{
		Commands: Commands{},
		Queries:  Queries{},
	}
}