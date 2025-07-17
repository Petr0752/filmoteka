package service

import "filmoteka/internal/model"

type ActorRepository interface {
	Create(a *model.Actor) (int64, error)
	Update(a *model.Actor) error
	Delete(id int64) error
	List() ([]model.Actor, error)
	GetByID(id int64) (*model.Actor, error)
}

type MovieRepository interface {
	Create(m *model.Movie) (int64, error)
	Update(m *model.Movie) error
	Delete(id int64) error
	List(sort string) ([]model.Movie, error)
	Search(query string) ([]model.Movie, error)
	FindByActorID(actorID int64) ([]model.Movie, error)
	AddActorToMovie(movieID, actorID int64) error
}
