package routes

import (
	"movie-festival-app/controllers"
	"movie-festival-app/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")

	admin := api.Group("/admin")
	admin.Use(middleware.AuthMiddleware())
	admin.POST("/movies", controllers.CreateMovie)
	admin.PUT("/movies/:id", controllers.UpdateMovie)
	admin.GET("/movies/most-viewed", controllers.GetMostViewedMovie)
	admin.GET("/genres/most-viewed", controllers.GetMostViewedGenre)

	api.GET("/movies", controllers.ListMovies)
	api.GET("/movies/search", controllers.SearchMovies)
	api.POST("/movies/:id/view", controllers.TrackMovieView)

	authUser := api.Group("/user")
	authUser.Use(middleware.AuthMiddleware())
	authUser.POST("/movies/:id/vote", controllers.VoteMovie)
	authUser.DELETE("/movies/:id/vote", controllers.UnvoteMovie)
	authUser.GET("/movies/votes", controllers.ListVotedMovies)

	api.POST("/register", controllers.RegisterUser)
	api.POST("/login", controllers.LoginUser)
}
