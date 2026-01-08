package routes

import (
	controller "github.com/Niteesh-Kulhari/StreamMovies/Server/StreamMovieServer/controllers"
	"github.com/Niteesh-Kulhari/StreamMovies/Server/StreamMovieServer/middleware"
	"github.com/gin-gonic/gin"
)

func SetupProtectedRoutes(router *gin.Engine) {
	router.Use(middleware.AuthMiddleware())

	// Protected Routes
	router.GET("/movie/:imdb_id", controller.GetMovie())
	router.POST("/addmovie", controller.AddMovie())

}
