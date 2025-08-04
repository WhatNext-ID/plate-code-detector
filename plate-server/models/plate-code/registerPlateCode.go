package platecode

import (
	"plate-server/models"

	"github.com/google/uuid"
)

type RegisterPlateCode struct {
	IdRegisterCode uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id_register"`
	RegisterCode   string    `gorm:"not null" valid:"required~Harap masukkan kode wilayah kendaraan" json:"register_code"`
	RegisterCity   string    `gorm:"not null" valid:"required~Harap masukkan wilayah kendaraan" json:"register_city"`
	Note           *string   `json:"note"`
	CodePosition   *int      `json:"code_position"`
	IdRegionCode   uuid.UUID `gorm:"type:uuid" json:"id_region"`
	models.GormModel
}
