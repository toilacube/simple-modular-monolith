package movies

import "github.com/gin-gonic/gin"

type MoviesHandler struct {
	service *MoviesService
}

func NewMoviesHandler(service *MoviesService) *MoviesHandler {
	return &MoviesHandler{
		service: service,
	}
}

// @Summary Create a new movie
// @Description Create a new movie associated with the authenticated member
// @Accept  json
// @Produce  json
// @Param   movie  body   CreateMovieDTO  true  "Create Movie"
// @Success 200 {object} MovieDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security Bearer token
// @Router /movies [post]
func (h *MoviesHandler) CreateMovie(c *gin.Context) {
	var createMovieDTO CreateMovieDTO
	if err := c.ShouldBindJSON(&createMovieDTO); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	createdMovie, err := h.service.CreateMovie(createMovieDTO, c.GetString("member_id"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, createdMovie)

}

// @Summary Get all movies by creator
// @Description Retrieve all movies created by the authenticated member
// @Produce  json
// @Success 200 {array} MovieDTO
// @Failure 500 {object} map[string]string
// @Security Bearer token
// @Router /movies [get]
func (h *MoviesHandler) GetMoviesByCreator(c *gin.Context) {
	creatorID := c.GetString("member_id")
	movies, err := h.service.GetMoviesByCreator(creatorID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, movies)
}
