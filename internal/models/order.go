package models

import "time"

type Order struct {
	ID        int
	UserID    int
	Total     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
