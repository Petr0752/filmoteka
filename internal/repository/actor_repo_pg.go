package repository

import (
	"database/sql"
	"filmoteka/internal/model"
)

type ActorPG struct {
	db *sql.DB
}

func NewActorPG(db *sql.DB) *ActorPG {
	return &ActorPG{db: db}
}

func (r *ActorPG) Create(a *model.Actor) (int64, error) {
	err := r.db.QueryRow(
		`INSERT INTO actors (name, gender, birth_date)
		 VALUES ($1,$2,$3) RETURNING id`,
		a.Name, a.Gender, a.BirthDate,
	).Scan(&a.ID)
	return a.ID, err
}

func (r *ActorPG) Update(a *model.Actor) error {
	_, err := r.db.Exec(
		`UPDATE actors SET name=$1, gender=$2, birth_date=$3, updated_at=now()
		 WHERE id=$4`,
		a.Name, a.Gender, a.BirthDate, a.ID)
	return err
}

func (r *ActorPG) Delete(id int64) error {
	_, err := r.db.Exec(`DELETE FROM actors WHERE id=$1`, id)
	return err
}

func (r *ActorPG) List() ([]model.Actor, error) {
	rows, err := r.db.Query(`SELECT id, name, gender, birth_date FROM actors`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []model.Actor
	for rows.Next() {
		var a model.Actor
		if err = rows.Scan(&a.ID, &a.Name, &a.Gender, &a.BirthDate); err != nil {
			return nil, err
		}
		res = append(res, a)
	}
	return res, rows.Err()
}

func (r *ActorPG) GetByID(id int64) (*model.Actor, error) {
	var a model.Actor
	err := r.db.QueryRow(
		`SELECT id, name, gender, birth_date FROM actors WHERE id=$1`, id,
	).Scan(&a.ID, &a.Name, &a.Gender, &a.BirthDate)
	if err != nil {
		return nil, err
	}
	return &a, nil
}
