package user

import (
	"server/models"

	"github.com/google/uuid"
)

type UserRole struct {
	IdRole   uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id_role"`
	RoleName string    `gorm:"not null" json:"role_name"`
	Users    []User    `gorm:"foreignKey:IdRole;references:IdRole"`
	models.GormModel
}
