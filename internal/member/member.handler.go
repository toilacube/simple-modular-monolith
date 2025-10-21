package member

import "github.com/gin-gonic/gin"

type MemberHandler struct {
	service *Service
}

func NewMemberHandler(service *Service) *MemberHandler {
	return &MemberHandler{service: service}
}

func (h *MemberHandler) Register(c *gin.Context) {
	var registerDTO RegisterDTO
	if err := c.ShouldBindJSON(&registerDTO); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Register(registerDTO); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "User registered successfully",
	})
}

func (h *MemberHandler) Login(c *gin.Context) {
	var loginDTO LoginDTO
	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := h.service.Login(loginDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
