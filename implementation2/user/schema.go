package user

import "time"

// Users : name of the table, generated from the config
type Users struct {
	Name string
	Roll int64
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
