package repository

import (
	"database/sql"
	"filmoteka/internal/model"
	"github.com/Masterminds/squirrel"
)

type ActorRepository struct {
	db *sql.DB
}

func NewActorRepository(db *sql.DB) *ActorRepository {
	return &ActorRepository{db: db}
}

func (r *ActorRepository) Create(a *model.Actor) (int64, error) {
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	query, args, err := psql.Insert("actors").
		Columns("name", "gender", "birth_date").
		Values(a.Name, a.Gender, a.BirthDate).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, err
	}
	return a.ID, r.db.QueryRow(query, args...).Scan(&a.ID)
}

func (r *ActorRepository) Update(a *model.Actor) error {
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	query, args, err := psql.Update("actors").
		Set("name", a.Name).
		Set("gender", a.Gender).
		Set("birth_date", a.BirthDate).
		Set("updated_at", squirrel.Expr("now()")).
		Where(squirrel.Eq{"id": a.ID}).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Exec(query, args...)
	return err
}

func (r *ActorRepository) Delete(id int64) error {
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	query, args, err := psql.Delete("actors").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Exec(query, args...)
	return err
}

func (r *ActorRepository) List() ([]model.Actor, error) {
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	query, args, err := psql.Select("id", "name", "gender", "birth_date").
		From("actors").
		ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := r.db.Query(query, args...)
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

func (r *ActorRepository) GetByID(id int64) (*model.Actor, error) {
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	query, args, err := psql.Select("id", "name", "gender", "birth_date").
		From("actors").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}
	var a model.Actor
	err = r.db.QueryRow(query, args...).Scan(&a.ID, &a.Name, &a.Gender, &a.BirthDate)
	if err != nil {
		return nil, err
	}
	return &a, nil
}
