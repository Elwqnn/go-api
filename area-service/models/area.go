package models

import "time"

type Area struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	UserID    uint       `json:"user_id"`   // Reference to the user who created this Area
	Name      string     `json:"name"`      // User-friendly name
	IsActive  bool       `json:"is_active"` // Whether this Area is enabled
	Actions   []Action   `json:"actions"`   // Associated actions
	Reactions []Reaction `json:"reactions"` // Associated reactions
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}
