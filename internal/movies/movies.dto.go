package movies

import "time"

type CreateMovieDTO struct {
	Name  string `json:"name" binding:"required"`
	Star  int    `json:"star" binding:"min=1,max=5"`
	Actor string `json:"actor" binding:"required"`
}

type MovieDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Star      int       `json:"star"`
	Actor     string    `json:"actor"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy string    `json:"created_by"`
}
