package app

import (
	"auth/internal/configs"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type App struct {
	Server  http.Server
	Context context.Context
	Cancel  context.CancelCauseFunc
	Config  *configs.Config
	Router  *gin.Engine
	//Logger  *logs.Logger
}

func NewApp() *App {
	return &App{
		Router: gin.New(),
		//Logger: logs.New(os.Stdout, logs.DEBUG),
		Config: configs.Load(),
	}
}

func (app *App) Run() {
	app.StartServer()
	app.GracefulShutdown()
}

func (app *App) StartServer() {
	app.Server = http.Server{
		Addr:    ":" + app.Config.Port,
		Handler: app.Router,
	}

	go func() {
		//app.Logger.Info("Start HTTP server on port %s", app.Config.Port)
		if err := app.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			//app.Logger.Error("Server ListenAndServe: %v", err)
		}
	}()
}

func (app *App) GracefulShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	//app.Logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.Server.Shutdown(ctx); err != nil {
		//app.Logger.Error("Server Shutdown: %v", err)
	} else {
		//app.Logger.Info("Server exited properly")
	}
}
