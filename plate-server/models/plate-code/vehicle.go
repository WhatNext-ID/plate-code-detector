package platecode

import (
	"server/models"
	"server/utils"

	"github.com/google/uuid"
)

type VehicleType struct {
	IdVehicleType   uuid.UUID         `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id_vehicle_type"`
	VehicleType     string            `gorm:"not null" valid:"required~Harap tambahkan jenis kendaraan" json:"vehicle_type"`
	VehicleCategory []VehicleCategory `gorm:"foreignKey:IdVehicleType;references:IdVehicleType"`
	models.GormModel
}

type VehicleEngine struct {
	IdVehicleEngine   uuid.UUID         `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id_vehicle_engine"`
	VehicleEngineType string            `gorm:"not null" valid:"required~Harap tambahkan jenis mesin kendaraan" json:"vehicle_engine_type"`
	VehicleCategory   []VehicleCategory `gorm:"foreignKey:IdVehicleEngine;references:IdVehicleEngine"`
	models.GormModel
}

type VehicleCategory struct {
	IdVehicleCategory uuid.UUID         `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id_vehicle_cat"`
	IdVehicleType     uuid.UUID         `gorm:"type:uuid" json:"id_type"`
	IdVehicleEngine   uuid.UUID         `gorm:"type:uuid" json:"id_engine"`
	ColorCriteria     utils.StringArray `gorm:"type:text[];not null" valid:"required~Harap tambahkan kriteria warna jenis kendaraan" json:"color_criteria"`
	models.GormModel
}
