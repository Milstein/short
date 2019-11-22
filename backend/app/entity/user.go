package entity

import "time"

// User represents an user of Short
type User struct {
	ID             string
	Name           string
	Email          string
	LastSignedInAt *time.Time
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}
