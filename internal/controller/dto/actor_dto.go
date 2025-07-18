package dto

import (
	"filmoteka/internal/model"
	"time"
)

type ActorDTO struct {
	ID        int64     `json:"id,omitempty"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	BirthDate time.Time `json:"birth_date"`
}

func ActorDTOToModel(dto *ActorDTO) *model.Actor {
	return &model.Actor{
		ID:        dto.ID,
		Name:      dto.Name,
		Gender:    dto.Gender,
		BirthDate: dto.BirthDate,
	}
}
