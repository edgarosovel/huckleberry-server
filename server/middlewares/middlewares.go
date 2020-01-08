package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"huckleberry.app/server/auth"
	"huckleberry.app/server/utils/formaterror"
)

func SetHeaders() gin.HandlerFunc {
	config := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Origin", "Content-Length", "Content-Type", "Accept"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	config.AllowAllOrigins = true

	return cors.New(config)
}

func Authentication() gin.HandlerFunc {
	return func(context *gin.Context) {
		usernameInToken, err := auth.IsTokenValid(context.Request)
		if err != nil {
			formattedError := formaterror.FormatError(http.StatusUnauthorized)
			context.JSON(http.StatusUnauthorized, formattedError)
			context.Abort()
			return
		}

		usernameInRoute := context.Param("username")
		if usernameInRoute != usernameInToken {
			formattedError := formaterror.FormatError(http.StatusForbidden)
			context.JSON(http.StatusForbidden, formattedError)
			context.Abort()
			return
		}
		context.Next()
	}
}
