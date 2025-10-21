package movies

import (
	"tutorial/pkg/model"

	"gorm.io/gorm"
)

type MoviesRepository struct {
	db *gorm.DB
}

func NewMoviesRepository(db *gorm.DB) *MoviesRepository {
	return &MoviesRepository{
		db: db,
	}
}

func (r *MoviesRepository) CreateMovie(movie model.Movie) (model.Movie, error) {
	if err := r.db.Create(&movie).Error; err != nil {
		return model.Movie{}, err
	}
	return movie, nil
}

func (r *MoviesRepository) GetMoviesByCreator(creatorID string) ([]model.Movie, error) {
	var movies []model.Movie
	err := r.db.Where("created_by = ?", creatorID).Find(&movies).Error
	if err != nil {
		return nil, err
	}
	return movies, nil
}
