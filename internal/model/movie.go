package model

import "time"

type Movie struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ReleaseDate time.Time `json:"release_date"`
	Rating      float32   `json:"rating"`
	Actors      []Actor   `json:"actors,omitempty"`
}
