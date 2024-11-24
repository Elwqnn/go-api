package models

import "time"

type Action struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	AreaID       uint      `json:"area_id"`       // reference to the associated Area
	Type         string    `json:"type"`          // e.g., "send_email", "post_webhook"
	Config       string    `json:"config"`        // JSON-encoded configuration
	IsSequential bool      `json:"is_sequential"` // if true, must wait for previous Action/Reaction
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
