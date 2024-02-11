package router

import (
	eventcontroller "KUNoti/internal/controller/event"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AppRouter struct {
	// notiController *NotiController
	// userController *UsertController
	// tagController *TagController
	eventController *eventcontroller.EventController
}

func (a AppRouter) InitEndpoints(r *gin.RouterGroup) {
	appGroup := r.Group("/api")
	a.eventController.InitEndpoints(appGroup)
}

func NewAppRouter(db *pgxpool.Pool) *AppRouter {
	return &AppRouter{
		eventController: eventcontroller.NewEventController(db),
	}
}
