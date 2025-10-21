package main

import (
	"tutorial/internal/app"
	"tutorial/internal/middleware"
	"tutorial/internal/movies"

	"github.com/gin-gonic/gin"
)

func main() {
	app, err := app.NewAppContainer()

	if err != nil {
		panic(err)
	}

	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.LoggerMiddleware())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})

	v1 := r.Group("/api/v1")
	{
		movieGroup := v1.Group("/movies")
		{
			movies.MovieProtectedRoutes(movieGroup, app.MoviesContainer)
		}
	}

	addr := ":" + app.Config.Server.MoviePort

	r.Run(addr)
}
