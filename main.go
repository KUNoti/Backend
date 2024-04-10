package main

import (
	"KUNoti/internal/router"
	"KUNoti/pkg/middleware"
	"context"
	"errors"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
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
		middleware.TimeoutMiddleware(),
	)

	firebaseApp, err := initializeFirebaseApp()
	if err != nil {
		log.Fatal(err)
	}

	routerGroup := r.Group("")
	router := router.NewAppRouter(db, firebaseApp)
	router.InitEndpoints(routerGroup)

	port := viper.GetString("SERVER_PORT")
	if port == "" {
		port = "8000"
	}
	log.Println("Server started on http://localhost:", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func initializeFirebaseApp() (*firebase.App, error) {
	opt := option.WithCredentialsFile("/Users/premkul/Documents/Workspace/final-project/Backend/ServiceAccount.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	return app, nil
}
