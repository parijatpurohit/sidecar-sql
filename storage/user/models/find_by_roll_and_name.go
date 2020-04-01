package models

import "time"


// User_FindByRollAndNameQuery ...
type User_FindByRollAndNameQuery struct {
	Roll      int64      `json:"Roll,omitempty"`
	DeletedAt *time.Time `json:"DeletedAt,omitempty"`
	UpdatedAt *time.Time `json:"UpdatedAt,omitempty"`
	Name      []string   `json:"Name,omitempty"`
}
