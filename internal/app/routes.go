package app

import "github.com/adarsh-devappsys/rest-with-protobuf-logs/internal/app/middleware"

func SetupRoutes(app *App) {
	middleware.MiddlewareWrapper(app.Router)

}
