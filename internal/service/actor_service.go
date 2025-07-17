package service

import (
	"errors"
	"filmoteka/internal/model"
	"filmoteka/internal/repository"
)

type ActorService struct {
	repo repository.ActorRepository
}

func NewActorService(r repository.ActorRepository) *ActorService { return &ActorService{repo: r} }

func (s *ActorService) Add(a *model.Actor) (int64, error) {
	if len(a.Name) == 0 {
		return 0, errors.New("name required")
	}
	return s.repo.Create(a)
}

func (s *ActorService) Update(a *model.Actor) error { return s.repo.Update(a) }
func (s *ActorService) Delete(id int64) error       { return s.repo.Delete(id) }

func (s *ActorService) ListWithMovies(mr repository.MovieRepository) ([]model.Actor, error) {
	actors, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	for i := range actors {
		movies, err := mr.Search(actors[i].Name)
		if err != nil {
			return nil, err
		}
		actors[i].Movies = movies
	}
	return actors, nil
}

func (s *ActorService) Get(id int64, mr repository.MovieRepository) (*model.Actor, error) {
	actor, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	movies, err := mr.FindByActorID(actor.ID)
	if err != nil {
		return nil, err
	}
	actor.Movies = movies

	return actor, nil
}
