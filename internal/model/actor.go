package model

import "time"

type Actor struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	BirthDate time.Time `json:"birth_date"`
	Movies    []Movie   `json:"movies,omitempty"`
}
