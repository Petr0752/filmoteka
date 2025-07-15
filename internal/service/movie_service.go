package service

import (
	"errors"
	"filmoteka/internal/model"
	"filmoteka/internal/repository"
)

type MovieService struct {
	repo repository.MovieRepository
}

func NewMovieService(r repository.MovieRepository) *MovieService {
	return &MovieService{repo: r}
}

func (s *MovieService) Add(m *model.Movie) (int64, error) {
	if l := len(m.Title); l == 0 || l > 150 {
		return 0, errors.New("title length must be 1-150")
	}
	if m.Rating < 0 || m.Rating > 10 {
		return 0, errors.New("rating must be 0-10")
	}
	return s.repo.Create(m)
}

func (s *MovieService) Update(m *model.Movie) error {
	return s.repo.Update(m)
}

func (s *MovieService) Delete(id int64) error {
	return s.repo.Delete(id)
}

func (s *MovieService) List(sort string) ([]model.Movie, error) { return s.repo.List(sort) }
func (s *MovieService) Search(q string) ([]model.Movie, error)  { return s.repo.Search(q) }
