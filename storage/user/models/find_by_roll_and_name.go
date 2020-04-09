package models


// User_FindByRollAndNameQuery ...
type User_FindByRollAndNameQuery struct {
	Name      []string   `json:"name,omitempty"`
	Roll      int64      `json:"roll,omitempty"`
}
