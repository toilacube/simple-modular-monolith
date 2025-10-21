package movies

import (
	"tutorial/internal/middleware"

	"github.com/gin-gonic/gin"
)

func MovieProtectedRoutes(r *gin.RouterGroup, container *MoviesContainer) {
	r.Use(middleware.AuthMiddleware())
	r.POST("", container.Handler.CreateMovie)
	r.GET("", container.Handler.GetMoviesByCreator)
}
