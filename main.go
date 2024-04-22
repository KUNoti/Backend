package main

import (
	"KUNoti/internal/router"
	"KUNoti/sqlc"
	"context"
	"errors"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/robfig/cron"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"time"
)

var appEnv string

func main() {

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("can't read from env")
		return
	}

	appEnv = viper.GetString("APP_ENV")
	if appEnv == "" {
		appEnv = "development"
	}

	r := gin.New()

	var db *pgxpool.Pool
	connStr := viper.GetString("PSQL_URL")
	if connStr == "" {
		panic(errors.New("no connection string"))
	}
	db, err = pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}

	var version string

	if err := db.QueryRow(context.Background(), "select version()").Scan(&version); err != nil {
		panic(err)
	}
	if appEnv == "development" {
		log.Println("connStr: ", connStr)
		log.Printf("Database: %s\n\n", version)
	}

	r.Use(
		gin.Logger(),
		gin.Recovery(),
		//middleware.TimeoutMiddleware(),
	)

	firebaseApp, err := initializeFirebaseApp()
	if err != nil {
		log.Fatal(err)
	}

	routerGroup := r.Group("")
	rout := router.NewAppRouter(db, firebaseApp)
	rout.InitEndpoints(routerGroup)

	port := viper.GetString("SERVER_PORT")
	if port == "" {
		port = "8000"
	}

	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Println(err)
	}
	cronJob := cron.NewWithLocation(loc) // Use NewWithLocation to create a cron instance with the specified location

	err = cronJob.AddFunc("0 */5 * * * *", func() { // Modified cron expression to run every 5 min
		log.Println("Every 5 min")
		queries := sqlc.New(db)
		events, err := queries.FindEventByTime(context.Background())
		if err != nil {
			log.Println(err)
			return
		}
		if len(events) != 0 {
			for i := range events {
				jsonEvent, err := json.Marshal(events[i])
				if err != nil {
					log.Println(err)
					return
				}

				users, err := queries.FindUserThatRegis(context.Background(), events[i].ID)
				if err != nil {
					log.Println(err)
					return
				}
				var userTokens []string
				if len(users) != 0 {
					for _, user := range users {
						userTokens = append(userTokens, user.Token)
					}
				}

				rout.EventController.Notification(context.Background(), userTokens, events[i], jsonEvent)
				_, err = queries.UpdateEventNoti(context.Background(), events[i].ID)
				if err != nil {
					log.Println(err)
					return
				}
			}
		}
	})
	if err != nil {
		log.Println(err)
		return
	}

	cronJob.Start()

	log.Println("Server started on http://localhost:", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func initializeFirebaseApp() (*firebase.App, error) {
	opt := option.WithCredentialsFile("ServiceAccount.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	return app, nil
}
