package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"tutorial/internal/app"
	"tutorial/internal/middleware"
	"tutorial/internal/movies"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "tutorial/docs/movie"
)

// @title           Movie Service API
// @version         1.0

// @host localhost
// @BasePath /v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Format: "Bearer {token}"
func main() {
	app, err := app.NewAppContainer()
	if err != nil {
		panic(err)
	}

	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.LoggerMiddleware())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	go func() {
		log.Println("Starting server on " + addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // block
	log.Println("Shutdown movie Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Unknow error, forced to shutdown movie server:", err)
	}

	log.Println("Server exiting")
}
