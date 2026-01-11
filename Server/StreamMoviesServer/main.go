package main

import (
	"fmt"

<<<<<<< HEAD
	"github.com/Niteesh-Kulhari/StreamMovies/Server/StreamMovieServer/routes"
=======
	controller "github.com/Niteesh-Kulhari/StreamMovies/Server/StreamMovieServer/controllers"

>>>>>>> 8de748d (Basic routes)
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello from Server")
	})
	//

	routes.SetupUnProtectedRoutes(router)
	routes.SetupProtectedRoutes(router)

	router.GET("/movies", controller.GetMovies())

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server", err)
	}
}
