package dto

import (
	"errors"
	"strings"
)

func ValidateActorDTO(a *ActorDTO) error {
	if len(strings.TrimSpace(a.Name)) < 3 {
		return errors.New("name must be at least 3 characters")
	}
	if a.Gender != "male" && a.Gender != "female" {
		return errors.New("gender must be 'male' or 'female'")
	}
	return nil
}

func ValidateMovieDTO(m *MovieDTO) error {
	if len(m.Title) < 2 {
		return errors.New("title must be at least 2 characters")
	}
	if m.Rating < 0 || m.Rating > 10 {
		return errors.New("rating must be between 0 and 10")
	}
	return nil
}
