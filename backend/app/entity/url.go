package entity

import "time"

// URL represents a mapping between an alias and a long link
type URL struct {
	Alias       string
	OriginalURL string
	ExpireAt    *time.Time
	CreatedBy   *User
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
