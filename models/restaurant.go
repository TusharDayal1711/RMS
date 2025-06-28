package models

import (
	"github.com/google/uuid"
	"time"
)

type Restaurant struct {
	ID         uuid.UUID  `db:"id" json:"id"`
	Name       string     `db:"name" json:"name"`
	Address    string     `db:"address" json:"address"`
	Longitude  float64    `db:"longitude" json:"longitude"`
	Latitude   float64    `db:"latitude" json:"latitude"`
	CreatedBy  uuid.UUID  `db:"created_by" json:"created_by"`
	CreatedAt  time.Time  `db:"created_at" json:"created_at"`
	ArchivedAt *time.Time `db:"archived_at" json:"archived_at,omitempty"` //can we omit in json if this is null
}

type RestaurantReq struct {
	ID        string  `json:"id" db:"id"`
	Name      string  `json:"name" db:"name"`
	Address   string  `json:"address" db:"address"`
	Longitude float64 `db:"longitude" json:"longitude"`
	Latitude  float64 `db:"latitude" json:"latitude"`
}
