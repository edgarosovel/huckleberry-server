package routers

import (
	"github.com/gin-gonic/gin"
	"huckleberry.app/server/controllers"
	"huckleberry.app/server/middlewares"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()

	router.Use(gin.Recovery()) // Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(middlewares.SetHeaders())

	api := router.Group("/api/v1")

	// no authentication endpoints
	{
		api.POST("login", controllers.Login)
		api.POST("users", controllers.CreateUser)
		api.GET("usernames/:username", controllers.UsernameExists)
		api.GET("emails/:email", controllers.EmailExists)
	}

	// basic authentication endpoints
	{
		basicAuth := api.Group("/")
		basicAuth.Use(middlewares.Authentication())
		{
			basicAuth.POST("users/:username/bookmarks", controllers.CreateBookmark)
			basicAuth.GET("users/:username/bookmarks", controllers.FindBookmarksByUsername)
			basicAuth.DELETE("users/:username/bookmarks/:id", controllers.DeleteBookmark)

			basicAuth.POST("users/:username/shares", controllers.CreateShare)
			basicAuth.GET("users/:username/shares", controllers.FindSharesByUsername)
			basicAuth.DELETE("users/:username/shares/:id", controllers.DeleteShare)

			basicAuth.GET("users/:username/notifications", controllers.FindNotificationsByUsername)
			basicAuth.PATCH("users/:username/notifications/:id", controllers.NotificationResponse)
			// basicAuth.POST("users/:username/bookmarks", controllers.CreateBookmark)
			// basicAuth.GET("/logout", logoutHandler)
		}
	}

	// admin authentication endpoints
	// {
	// 	adminAuth := api.Group("/admin")
	// 	adminAuth.Use(AuthenticationRequired("admin"))
	// 	{
	// 		adminAuth.GET("/message/:msg", adminMessageHandler)
	// 	}
	// }

	//Users routes
	// s.Router.HandleFunc("/api/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	// s.Router.HandleFunc("/api/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	// s.Router.HandleFunc("/api/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	// s.Router.HandleFunc("/api/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//User's bookmark routes
	// s.Router.HandleFunc("/api/users/{id}/bookmarks", middlewares.SetMiddlewareJSON(s.CreateBookmark)).Methods("POST")
	// s.Router.HandleFunc("/api/users/{id}/bookmarks", middlewares.SetMiddlewareJSON(s.GetBookmarks)).Methods("GET")
	// s.Router.HandleFunc("/api/users/{user_id}/bookmarks/{bookmark_id}", middlewares.SetMiddlewareAuthentication(s.DeleteBookmark)).Methods("DELETE")

	return router
}
