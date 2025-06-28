package models

import (
	"github.com/google/uuid"
	"time"
)

type Address struct {
	ID         uuid.UUID  `db:"id" json:"id"`
	UserID     uuid.UUID  `db:"user_id" json:"user_id"`
	Address    string     `db:"address" json:"address"`
	Longitude  float64    `db:"longitude" json:"longitude"`
	Latitude   float64    `db:"latitude" json:"latitude"`
	CreatedAt  time.Time  `db:"created_at" json:"created_at"`
	ArchivedAt *time.Time `db:"archived_at" json:"archived_at,omitempty"`
}

type AddCoord struct {
	Longitude float64 `db:"longitude"`
	Latitude  float64 `db:"latitude"`
}

type AddressReq struct {
	Address   string  `json:"address"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
