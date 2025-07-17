package repository

import (
	"database/sql"
	"filmoteka/internal/model"
	"fmt"
	"strings"
)

type MoviePG struct{ db *sql.DB }

func NewMoviePG(db *sql.DB) *MoviePG { return &MoviePG{db: db} }

func (r *MoviePG) Create(m *model.Movie) (int64, error) {
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

func (r *MoviePG) Update(m *model.Movie) error {
	_, err := r.db.Exec(
		`UPDATE movies SET title=$1, description=$2, release_date=$3, rating=$4,
		       updated_at=now() WHERE id=$5`,
		m.Title, m.Description, m.ReleaseDate, m.Rating, m.ID)
	return err
}

func (r *MoviePG) Delete(id int64) error {
	_, err := r.db.Exec(`DELETE FROM movies WHERE id=$1`, id)
	return err
}

var allowedSort = map[string]string{
	"title":        "title",
	"rating":       "rating",
	"release_date": "release_date",
}

func (r *MoviePG) List(sort string) ([]model.Movie, error) {
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

func (r *MoviePG) Search(q string) ([]model.Movie, error) {
	p := "%" + strings.ToLower(q) + "%"
	rows, err := r.db.Query(`
	  SELECT DISTINCT m.id, m.title, m.description, m.release_date, m.rating
	  FROM movies m
	  LEFT JOIN movie_actors ma ON m.id = ma.movie_id
	  LEFT JOIN actors a ON a.id = ma.actor_id
	  WHERE lower(m.title) LIKE $1 OR lower(a.name) LIKE $1`, p)
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

func (r *MoviePG) FindByActorID(actorID int64) ([]model.Movie, error) {
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

func (r *MoviePG) AddActorToMovie(movieID, actorID int64) error {
	_, err := r.db.Exec(`
		INSERT INTO movie_actors (movie_id, actor_id)
		VALUES ($1, $2)
		ON CONFLICT DO NOTHING
	`, movieID, actorID)
	return err
}
