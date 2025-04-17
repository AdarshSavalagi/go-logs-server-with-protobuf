package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		// AllowedOrigins defines the list of origins allowed to access the server.
		// "*" allows any origin, but it is recommended to specify domains for security reasons.
		AllowOrigins: []string{
			"https://fintracker.uatdemo.info",
			"http://localhost:3000",
		},
		// AllowMethods specifies the HTTP methods allowed for cross-origin requests.
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"},

		// AllowHeaders defines the headers that the client can send with requests.
		AllowHeaders: []string{
			"Origin", "Content-Length", "Content-Type", "Accept", "Authorization",
			"Idempotency-Key", "Accept-Encoding", "X-Request-Id", "X-CSRF-Token",
		},
		// AllowOriginFunc: func(origin string) bool {
		// 	return strings.HasPrefix(origin, "http://localhost")
		// },

		// ExposeHeaders defines the headers that are accessible from the client side.
		ExposeHeaders: []string{"Content-Length"},

		// AllowCredentials indicates whether credentials (cookies, authorization headers) can be included.
		AllowCredentials: true,

		// MaxAge specifies how long (in seconds) the results of a preflight request can be cached.
		MaxAge: 12 * time.Hour,

		// AllowFiles and AllowWebSockets determine whether file uploads or WebSocket connections are allowed.
		AllowFiles:      false,
		AllowWebSockets: false,
	})
}
