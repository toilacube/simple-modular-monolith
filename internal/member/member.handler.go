package member

import "github.com/gin-gonic/gin"

type MemberHandler struct {
	service *Service
}

func NewMemberHandler(service *Service) *MemberHandler {
	return &MemberHandler{service: service}
}

// @Summary Register a new member
// @Description Add a new member to the database
// @Accept  json
// @Produce  json
// @Param   member  body   RegisterDTO  true  "Register Member"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /member/register [post]
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

// @Summary Login a member
// @Description Authenticate a member and return a JWT token
// @Accept  json
// @Produce  json
// @Param   member  body   LoginDTO  true  "Login Member"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /member/login [post]
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
