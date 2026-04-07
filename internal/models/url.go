package models

import (
	"time"

	"gorm.io/gorm"
)

type URL struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	LongURL   string         `gorm:"type:text;not null" json:"long_url"`
	ShortCode string         `gorm:"uniqueIndex;not null" json:"short_code"`
	Clicks    int            `gorm:"default:0" json:"clicks"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
