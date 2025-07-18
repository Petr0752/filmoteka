package repository

import (
	"database/sql"
	"filmoteka/internal/model"
	"github.com/Masterminds/squirrel"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByUsername(username string) (*model.User, error) {
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	queryBuilder := psql.
		Select("id", "username", "password", "role").
		From("users").
		Where(squirrel.Eq{"username": username})

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var u model.User
	err = r.db.QueryRow(query, args...).
		Scan(&u.ID, &u.Username, &u.Password, &u.Role)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
