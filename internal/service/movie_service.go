package service

import (
	"filmoteka/internal/model"
)

type MovieService struct {
	repo MovieRepository
}

func NewMovieService(r MovieRepository) *MovieService {
	return &MovieService{repo: r}
}

func (s *MovieService) Add(movie *model.Movie) (int64, error) {
	return s.repo.Create(movie)
}

func (s *MovieService) Update(movie *model.Movie) error { return s.repo.Update(movie) }

func (s *MovieService) Delete(id int64) error { return s.repo.Delete(id) }

func (s *MovieService) List(sort string) ([]model.Movie, error) { return s.repo.List(sort) }

func (s *MovieService) Search(q string) ([]model.Movie, error) { return s.repo.Search(q) }

func (s *MovieService) AddActor(movieID, actorID int64) error {
	return s.repo.AddActorToMovie(movieID, actorID)
}
