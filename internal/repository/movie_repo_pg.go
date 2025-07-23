package repository

import (
	"database/sql"
	"filmoteka/internal/model"
	"github.com/Masterminds/squirrel"
)

type MovieRepository struct {
	db *sql.DB
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{db: db}
}

func (r *MovieRepository) Create(m *model.Movie) (int64, error) {
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	query, args, err := psql.Insert("movies").
		Columns("title", "description", "release_date", "rating").
		Values(m.Title, m.Description, m.ReleaseDate, m.Rating).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, err
	}
	return m.ID, r.db.QueryRow(query, args...).Scan(&m.ID)
}

func (r *MovieRepository) Update(m *model.Movie) error {
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	query, args, err := psql.
		Update("movies").
		Set("title", m.Title).
		Set("description", m.Description).
		Set("release_date", m.ReleaseDate).
		Set("rating", m.Rating).
		Set("updated_at", squirrel.Expr("now()")).
		Where(squirrel.Eq{"id": m.ID}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *MovieRepository) Delete(id int64) error {
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	query, args, err := psql.
		Delete("movies").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

var allowedSort = map[string]string{
	"title":        "title",
	"rating":       "rating",
	"release_date": "release_date",
}

func (r *MovieRepository) List(sort string) ([]model.Movie, error) {
	field, ok := allowedSort[sort]
	if !ok {
		field = "rating"
	}

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	queryBuilder := psql.
		Select("id", "title", "description", "release_date", "rating").
		From("movies").
		OrderBy(field + " DESC")

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []model.Movie
	for rows.Next() {
		var m model.Movie
		if err = rows.Scan(&m.ID, &m.Title, &m.Description, &m.ReleaseDate, &m.Rating); err != nil {
			return nil, err
		}
		res = append(res, m)
	}
	return res, rows.Err()
}

func (r *MovieRepository) Search(q string) ([]model.Movie, error) {
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	queryBuilder := psql.Select("m.id", "m.title", "m.description", "m.release_date", "m.rating").
		From("movies m").
		LeftJoin("movie_actors ma ON m.id = ma.movie_id").
		LeftJoin("actors a ON a.id = ma.actor_id")

	if q != "" {
		queryBuilder = queryBuilder.Where(
			squirrel.Or{
				squirrel.Like{"m.title": "%" + q + "%"},
				squirrel.Like{"m.description": "%" + q + "%"},
				squirrel.Like{"a.name": "%" + q + "%"},
			},
		)
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []model.Movie
	for rows.Next() {
		var m model.Movie
		if err = rows.Scan(&m.ID, &m.Title, &m.Description, &m.ReleaseDate, &m.Rating); err != nil {
			return nil, err
		}
		res = append(res, m)
	}
	return res, rows.Err()
}

func (r *MovieRepository) FindByActorID(actorID int64) ([]model.Movie, error) {
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	query, args, err := psql.
		Select("m.id", "m.title", "m.description", "m.release_date", "m.rating").
		From("movies m").
		Join("movie_actors ma ON ma.movie_id = m.id").
		Where(squirrel.Eq{"ma.actor_id": actorID}).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []model.Movie
	for rows.Next() {
		var m model.Movie
		if err := rows.Scan(&m.ID, &m.Title, &m.Description, &m.ReleaseDate, &m.Rating); err != nil {
			return nil, err
		}
		movies = append(movies, m)
	}
	return movies, nil
}

func (r *MovieRepository) AddActorToMovie(movieID, actorID int64) error {
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	query, args, err := psql.
		Insert("movie_actors").
		Columns("movie_id", "actor_id").
		Values(movieID, actorID).
		Suffix("ON CONFLICT DO NOTHING").
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}
