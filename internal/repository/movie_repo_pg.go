package repository

import (
	"database/sql"
	"filmoteka/internal/model"
	"fmt"
	"github.com/Masterminds/squirrel"
)

type MovieRepository struct {
	db *sql.DB
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{db: db}
}

func (r *MovieRepository) Create(m *model.Movie) (int64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	err = tx.QueryRow(
		`INSERT INTO movies (title, description, release_date, rating)
		 VALUES ($1,$2,$3,$4) RETURNING id`,
		m.Title, m.Description, m.ReleaseDate, m.Rating).Scan(&m.ID)
	if err != nil {
		return 0, err
	}

	for _, a := range m.Actors {
		if _, err = tx.Exec(
			`INSERT INTO movie_actors (movie_id, actor_id) VALUES ($1,$2)`,
			m.ID, a.ID); err != nil {
			return 0, err
		}
	}
	return m.ID, tx.Commit()
}

func (r *MovieRepository) Update(m *model.Movie) error {
	_, err := r.db.Exec(
		`UPDATE movies SET title=$1, description=$2, release_date=$3, rating=$4,
		       updated_at=now() WHERE id=$5`,
		m.Title, m.Description, m.ReleaseDate, m.Rating, m.ID)
	return err
}

func (r *MovieRepository) Delete(id int64) error {
	_, err := r.db.Exec(`DELETE FROM movies WHERE id=$1`, id)
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
		field = "rating" // default: по убыв. рейтинга
	}
	query := fmt.Sprintf(`SELECT id, title, description, release_date, rating
	                       FROM movies ORDER BY %s DESC`, field)

	rows, err := r.db.Query(query)
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
	rows, err := r.db.Query(`
		SELECT m.id, m.title, m.description, m.release_date, m.rating
		FROM movies m
		JOIN movie_actors ma ON ma.movie_id = m.id
		WHERE ma.actor_id = $1
	`, actorID)
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
	_, err := r.db.Exec(`
		INSERT INTO movie_actors (movie_id, actor_id)
		VALUES ($1, $2)
		ON CONFLICT DO NOTHING
	`, movieID, actorID)
	return err
}
