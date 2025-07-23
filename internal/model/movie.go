package model

import "time"

type Movie struct {
	ID          int64
	Title       string
	Description string
	ReleaseDate time.Time
	Rating      float32
	Actors      []Actor
}
