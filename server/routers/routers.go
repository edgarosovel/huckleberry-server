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
		api.GET("users/username/:username", controllers.UsernameExists)
		api.GET("users/email/:email", controllers.EmailExists)
	}

	// basic authentication endpoints
	// {
	// 	basicAuth := api.Group("/")
	// 	basicAuth.Use(AuthenticationRequired())
	// 	{
	// 		basicAuth.GET("/logout", logoutHandler)
	// 	}
	// }

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
