package models

import (
	"time"
)

type GormModel struct {
	CreatedAt *time.Time `gorm:"CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `gorm:"CURRENT_TIMESTAMP" json:"updated_at"`
}
