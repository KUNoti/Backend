package repository

import (
	"KUNoti/sqlc"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FirebaseRepository struct {
	DB      *pgxpool.Pool
	queries *sqlc.Queries
}

type CreateNoti struct {
	Body  string `json:"body"`
	Data  []byte `json:"data"`
	Title string `json:"title"`
	Token string `json:"token"`
}

func (fb FirebaseRepository) Create(ctx context.Context, param CreateNoti) error {
	_, err := fb.queries.CreateNotification(ctx, sqlc.CreateNotificationParams{
		Body:  param.Body,
		Data:  param.Data,
		Title: param.Title,
		Token: param.Token,
	})
	if err != nil {
		return err
	}
	return nil
}

type Notification struct {
	Title string
	Body  string
	Data  []byte
	Token string
}

func (fb FirebaseRepository) FindByToken(ctx context.Context, token string) ([]Notification, error) {
	notis, err := fb.queries.FindNotiByToken(ctx, token)
	if err != nil {
		return nil, err
	}
	returnNotis := make([]Notification, len(notis))
	for i, n := range notis {
		returnNotis[i] = Notification{
			Title: n.Title,
			Body:  n.Body,
			Data:  n.Data,
			Token: n.Token,
		}
	}
	return returnNotis, nil
}
