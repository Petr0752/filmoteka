package model

import "time"

type Actor struct {
	ID        int64
	Name      string
	Gender    string
	BirthDate time.Time
	Movies    []Movie
}
