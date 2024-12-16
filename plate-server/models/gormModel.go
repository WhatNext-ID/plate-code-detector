package models

import (
	"time"

	"gorm.io/gorm"
)

type GormModel struct {
	IdStatus  int        `gorm:"default:1" json:"id_status"`
	CreatedAt *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time `gorm:"null;default:null" json:"updated_at"`
}

// BeforeCreate sets the CreatedAt and UpdatedAt fields to the local time before a record is created.
func (m *GormModel) BeforeCreate(*gorm.DB) error {
	now := time.Now().Local()
	m.CreatedAt = &now
	return nil
}

// BeforeUpdate sets the UpdatedAt field to the local time before a record is updated.
func (m *GormModel) BeforeUpdate(*gorm.DB) error {
	now := time.Now().Local()
	m.UpdatedAt = &now
	return nil
}
