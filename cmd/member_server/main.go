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
	"tutorial/internal/member"
	"tutorial/internal/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "tutorial/docs/member"
)

// @title           Member Service API
// @version         1.0
// @description     This is the API documentation for the Member Service.

// @host http://localhost
// @BasePath /v1
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
		memberGroup := v1.Group("/member")
		{
			member.AuthRoutes(memberGroup, app.MemberContainer)

		} // Logger test endpoints group
		loggerGroup := v1.Group("/logger")
		{
			loggerGroup.GET("/get", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "Logger test GET endpoint",
				})
			})

			loggerGroup.POST("/post", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "Logger test POST endpoint",
				})
			})

			// logger for long response time endpoint
			loggerGroup.GET("/slow", func(c *gin.Context) {
				time.Sleep(1200 * time.Millisecond)
				c.JSON(200, gin.H{
					"message": "Logger test SLOW endpoint",
				})
			})

			loggerGroup.GET("/error", func(c *gin.Context) {
				c.JSON(500, gin.H{
					"message": "Logger test ERROR endpoint",
				})
			})
		}
	}

	addr := ":" + app.Config.Server.MemberPort

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
	log.Println("Shutdown member Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Unknow error, forced to shutdown member server:", err)
	}

	log.Println("Server exiting")
}
