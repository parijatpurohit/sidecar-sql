package models

import "time"

// User : name of the table, generated from the config
type User struct {
	UUID string
	Name string
	Roll int64
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
