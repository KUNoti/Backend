package main

import (
	userdomain "KUNoti/internal/user/domain/user"
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

var appEnv string

func testResponse(c *gin.Context) {
	c.String(http.StatusRequestTimeout, "timeout")
}

func timeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(1*time.Minute),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(testResponse),
	)
}

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
		timeoutMiddleware(),
	)

	r.GET("/", testResponse)
	//r.POST("/users/login", userdomain.LoginUser(context.Background()))
	test := userdomain.Test{
		ID: 3,
	}
	log.Println(test)

	port := viper.GetString("SERVER_PORT")
	if port == "" {
		port = "8000"
	}
	log.Println("Server started on http://localhost:", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
