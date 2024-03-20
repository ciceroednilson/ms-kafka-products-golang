package entities

import (
	"time"
)

type Product struct {
	Id          int       `db:"id"`
	Description string    `db:"description"`
	Price       float64   `db:"price"`
	Total       int16     `db:"total"`
	Created     time.Time `db:"created"`
}
