package app

import (
	logs_api "github.com/adarsh-devappsys/rest-with-protobuf-logs/internal/app/api/logs"
	"github.com/adarsh-devappsys/rest-with-protobuf-logs/internal/app/middleware"
)

func SetupRoutes(app *App) {
	middleware.MiddlewareWrapper(app.Router)

	logs_api.Routes(app.Router,app.KafkaWriter,app.Logger)
}
