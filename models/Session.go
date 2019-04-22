package models

import "time"

type Session struct {
	SessionID string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiresAt time.Time
}
