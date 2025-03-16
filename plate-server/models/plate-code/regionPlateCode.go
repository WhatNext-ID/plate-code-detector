package platecode

import (
	"server/models"

	"github.com/google/uuid"
)

type RegionPlateCode struct {
	IdRegionCode      uuid.UUID           `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id_region"`
	RegionCode        string              `gorm:"not null" valid:"required~Harap masukkan kode wilayah kendaraan" json:"region_code"`
	RegionArea        string              `gorm:"not null" valid:"required~Harap masukkan wilayah kendaraan" json:"region_area"`
	Note              *string             `json:"note"`
	RegisterPlateCode []RegisterPlateCode `gorm:"foreignKey:IdRegionCode;references:IdRegionCode"`
	models.GormModel
}
