package platecode

import (
	"server/models"

	"github.com/google/uuid"
)

type RegionPlateCode struct {
	IdRegionCode      uuid.UUID           `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id_region"`
	RegionCode        string              `gorm:"not null" valid:"required~Harap masukkan kode wilayah kendaraan" json:"region_code"`
	RegionArea        string              `gorm:"not null" valid:"required~Harap masukkan wilayah kendaraan" json:"region_area"`
	Note              string              `json:"note"`
	IdCodePosition    uuid.UUID           `gorm:"type:uuid" json:"id_code_position"`
	RegisterPlateCode []RegisterPlateCode `gorm:"foreignKey:IdRegionCode;references:IdRegionCode"`
	models.GormModel
}

type RegisterCodePosition struct {
	IdCodePosition  uuid.UUID         `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id_code_position"`
	CodePosition    string            `gorm:"not null" valid:"required~Harap masukkan posisi kode registrasi kendaraan" json:"code_position"`
	RegionPlateCode []RegionPlateCode `gorm:"foreignKey:IdCodePosition;references:IdCodePosition"`
	models.GormModel
}
