package movies

type MoviesService struct {
	repository *MoviesRepository
}

func NewMoviesService(repository *MoviesRepository) *MoviesService {
	return &MoviesService{
		repository: repository,
	}
}

func (s *MoviesService) CreateMovie(createMovieDTO CreateMovieDTO, creatorID string) (MovieDTO, error) {
	var movie = ConvertCreateMovieDTOToMovie(createMovieDTO, creatorID)
	createdMovie, err := s.repository.CreateMovie(movie)
	if err != nil {
		return MovieDTO{}, err
	}
	return ConvertMovieToMovieDTO(createdMovie), nil
}

func (s *MoviesService) GetMoviesByCreator(creatorID string) ([]MovieDTO, error) {
	movies, err := s.repository.GetMoviesByCreator(creatorID)
	if err != nil {
		return nil, err
	}

	var movieDTOs []MovieDTO
	for _, movie := range movies {
		movieDTOs = append(movieDTOs, ConvertMovieToMovieDTO(movie))
	}
	return movieDTOs, nil
}
