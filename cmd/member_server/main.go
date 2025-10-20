package main

import (
	"tutorial/internal/app"

	"github.com/gin-gonic/gin"
)

func main() {
	app, err := app.NewAppContainer()

	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})

	memberServerPrefix := "/member"

	v1 := r.Group("/api/v1")
	{
		v1.GET(memberServerPrefix, func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Get all members",
			})
		})
	}

	addr := ":" + app.Config.MemberServerPort

	r.Run(addr)
}
