package models

import (
	"github.com/google/uuid"
	"time"
)

type Session struct {
	ID           uuid.UUID  `db:"id" json:"id"`
	UserID       uuid.UUID  `db:"user_id" json:"user_id"`
	CreatedAt    time.Time  `db:"created_at" json:"created_at"`
	RefreshToken string     `db:"refresh_token" json:"refresh_token"`
	ExpireAt     *time.Time `db:"expire_at" json:"expire_at,omitempty"`
	ArchivedAt   *time.Time `db:"archived_at" json:"archived_at,omitempty"`
}
