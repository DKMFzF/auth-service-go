package app

import (
	"auth/internal/middleware"
	"auth/internal/routes"

	"github.com/gin-gonic/gin"
)

func Bootstrap() *App {

	app := NewApp()

	app.Router.Use(gin.Recovery(), middleware.ErrorHandler(), middleware.ErrorHandler())
	routes.SetupRoutes(app.Router)

	return app
}
