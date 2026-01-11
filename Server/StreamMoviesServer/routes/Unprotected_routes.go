package routes

import (
	controller "github.com/Niteesh-Kulhari/StreamMovies/Server/StreamMovieServer/controllers"
	"github.com/gin-gonic/gin"
)

func SetupUnProtectedRoutes(router *gin.Engine) {

	// Unprotected Routes
	router.GET("/movies", controller.GetMovies())
	router.POST("/register", controller.RegisterUser())
	router.POST("/login", controller.LoginUser())
	router.PATCH("/updatereview/:imdb_id", controller.AdminReviewUpdate())
}
