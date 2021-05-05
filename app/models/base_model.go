package models

import "time"

type BaseModel struct {
	ID        int64     `json:"id"`
	IsValid   int       `json:"is_valid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
