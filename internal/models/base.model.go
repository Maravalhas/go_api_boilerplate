package models

import "time"

type BaseModel struct {
	ID uint `gorm:"primarykey" json:"id"`
}

type BaseModelWithTimestamps struct {
	BaseModel
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
