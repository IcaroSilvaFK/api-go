package entities

import "time"

type Category struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	CreateAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

