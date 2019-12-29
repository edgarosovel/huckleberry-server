package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"huckleberry.app/server/auth"
)

func SetHeaders() gin.HandlerFunc {
	return cors.Default()

	// return func(context *gin.Context) {
	// 	context.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	// 	context.Writer.Header().Add("Access-Control-Max-Age", "10000")
	// 	context.Writer.Header().Add("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
	// 	context.Writer.Header().Add("Access-Control-Allow-Headers", "Authorization,Origin,Content-Type,Accept")
	// 	context.Writer.Header().Set("Content-Type", "application/json")
	// 	context.Next()
	// }
}

func SetMiddlewareAuthentication() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := auth.IsTokenValid(context.Request)
		if err != nil {
			// responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		context.Next()
	}
}
