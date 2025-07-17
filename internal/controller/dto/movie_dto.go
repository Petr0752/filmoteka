package dto

import (
	"filmoteka/internal/model"
	"time"
)

type MovieDTO struct {
	ID          int64     `json:"id,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ReleaseDate time.Time `json:"release_date"`
	Rating      float32   `json:"rating"`
}

func MovieDTOToModel(dto *MovieDTO) *model.Movie {
	return &model.Movie{
		ID:          dto.ID,
		Title:       dto.Title,
		Description: dto.Description,
		ReleaseDate: dto.ReleaseDate,
		Rating:      dto.Rating,
	}
}

func MovieModelToDTO(m *model.Movie) *MovieDTO {
	return &MovieDTO{
		ID:          m.ID,
		Title:       m.Title,
		Description: m.Description,
		ReleaseDate: m.ReleaseDate,
		Rating:      m.Rating,
	}
}
