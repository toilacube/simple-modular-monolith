package movies

import "gorm.io/gorm"

type MoviesContainer struct {
	Handler    *MoviesHandler
	Service    *MoviesService
	Repository *MoviesRepository
}

func NewMoviesContainer(db *gorm.DB) *MoviesContainer {
	repository := NewMoviesRepository(db)
	service := NewMoviesService(repository)
	handler := NewMoviesHandler(service)

	return &MoviesContainer{
		Handler:    handler,
		Service:    service,
		Repository: repository,
	}
}
