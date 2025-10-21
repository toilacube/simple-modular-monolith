package movies

import (
	"tutorial/pkg/model"
	"tutorial/pkg/utils"
)

func ConvertMovieToMovieDTO(movie model.Movie) MovieDTO {
	return MovieDTO{
		ID:        movie.ID,
		Name:      movie.Name,
		Star:      movie.Star,
		Actor:     movie.Actor,
		CreatedAt: movie.CreatedAt,
		UpdatedAt: movie.UpdatedAt,
		CreatedBy: movie.CreatedBy,
	}
}

func ConvertCreateMovieDTOToMovie(dto CreateMovieDTO, creatorID string) model.Movie {
	return model.Movie{
		ID:        utils.GetUUID(),
		Name:      dto.Name,
		Star:      dto.Star,
		Actor:     dto.Actor,
		CreatedBy: creatorID,
	}
}
