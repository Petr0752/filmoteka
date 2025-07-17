package repository

import "filmoteka/internal/model"

type MovieRepository interface {
	Create(m *model.Movie) (int64, error)
	Update(m *model.Movie) error
	Delete(id int64) error
	List(sort string) ([]model.Movie, error)
	Search(query string) ([]model.Movie, error)
	FindByActorID(actorID int64) ([]model.Movie, error)
	AddActorToMovie(movieID, actorID int64) error
}
