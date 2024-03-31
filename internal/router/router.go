package router

import (
	eventcontroller "KUNoti/internal/controller/event"
	usercontroller "KUNoti/internal/controller/user"
	fbService "KUNoti/service/firebaseService"
	firebase "firebase.google.com/go/v4"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AppRouter struct {
	eventController *eventcontroller.EventController
	userController  *usercontroller.UserController
}

func (a AppRouter) InitEndpoints(r *gin.RouterGroup) {
	appGroup := r.Group("/api")
	a.eventController.InitEndpoints(appGroup)
	a.userController.InitEndpoints(appGroup)
}

func NewAppRouter(db *pgxpool.Pool, firebaseApp *firebase.App) *AppRouter {

	firebaseService := fbService.NewFirebaseServiceClient(firebaseApp)

	return &AppRouter{
		eventController: eventcontroller.NewEventController(db, firebaseService),
		userController:  usercontroller.NewUserController(db),
	}
}
