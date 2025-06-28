package models

import (
	"github.com/google/uuid"
	"time"
)

type dishes struct {
	Name         string     `db:"name" json:"name"`
	Price        float32    `db:"price" json:"price"`
	RestaurantID uuid.UUID  `db:"restaurant_id" json:"restaurant_id"`
	CreatedBy    uuid.UUID  `db:"created_by" json:"created_by"`
	CreatedAt    time.Time  `db:"created_at" json:"created_at"`
	ArchivedAt   *time.Time `db:"archived_at" json:"archived_at,omitempty"`
}

type DishReq struct {
	Name         string  `json:"name" db:"name"`
	Price        float64 `json:"price" db:"price"`
	RestaurantID string  `json:"restaurant_id" db:"restaurant_id"`
}

type AllDishReq struct {
	ID           uuid.UUID `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	Price        float64   `db:"price" json:"price"`
	RestaurantID uuid.UUID `db:"restaurant_id" json:"restaurant_id"`
	CreatedBy    uuid.UUID `db:"created_by" json:"created_by"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}
