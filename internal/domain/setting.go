package domain

import "time"

// Setting represents a single configuration entry in the settings table.
type Setting struct {
	ID        string    `json:"id" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Key       string    `json:"key" gorm:"column:key;unique;not null"`
	Value     string    `json:"value" gorm:"column:value;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;not null;default:now()"`
}
