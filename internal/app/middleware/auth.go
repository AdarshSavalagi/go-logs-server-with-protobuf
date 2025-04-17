package middleware

import (
	"strings"

	"github.com/adarsh-devappsys/rest-with-protobuf-logs/internal/app/constants"
	"github.com/adarsh-devappsys/rest-with-protobuf-logs/internal/app/utils/auth"
	"github.com/adarsh-devappsys/rest-with-protobuf-logs/internal/utils/response"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the Authorization header
		authHeader := c.GetHeader(constants.Const.Headers.AuthToken)
		if authHeader == "" {
			// If no token is found, return unauthorized response
			response.RespondUnauthorized(c)
			c.Abort()
			return
		}

		// Remove "Bearer " prefix from the token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Initialize claims struct
		claims := &auth.Claims{}

		// Parse and validate JWT token, extracting claims
		claims, err := auth.GetJwtClaims(tokenString)
		if err != nil {
			// If token is invalid, return unauthorized response
			response.RespondUnauthorized(c)
			c.Abort()
			return
		}

		// Store user information in the Gin context
		c.Set(constants.Const.Context.UserID, claims.Sub)
		c.Set(constants.Const.Context.Username, claims.Email)

		// Proceed to the next middleware or handler
		c.Next()
	}
}
