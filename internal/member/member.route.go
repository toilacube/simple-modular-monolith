package member

import (
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup, container *MemberContainer) {
	r.POST("/register", container.Handler.Register)
	r.POST("/login", container.Handler.Login)
}
