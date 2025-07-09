package repository

import "filmoteka/internal/model"

type ActorRepository interface {
	Create(a *model.Actor) (int64, error)
	Update(a *model.Actor) error
	Delete(id int64) error
	List() ([]model.Actor, error)
}
