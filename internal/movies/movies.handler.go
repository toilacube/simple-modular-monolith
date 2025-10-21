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

func (h *MoviesHandler) GetMoviesByCreator(c *gin.Context) {
	creatorID := c.GetString("member_id")
	movies, err := h.service.GetMoviesByCreator(creatorID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, movies)
}
