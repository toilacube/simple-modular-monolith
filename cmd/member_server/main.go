package main

import (
	"time"
	"tutorial/internal/app"
	"tutorial/internal/member"
	"tutorial/internal/middleware"

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

	r.Run(addr)
}
