package models

import (
	"github.com/google/uuid"
)

type MStatusKendaraan struct {
	IdStatusKendaraan uuid.UUID              `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id_status_kendaraan"`
	StatusKendaraan   string                 `gorm:"not null" valid:"required~Harap tambahkan status kendaraan" json:"status_kendaraan"`
	Keterangan        string                 `gorm:"not null" valid:"required~Harap tambahkan keterangan status kendaraan" json:"keterangan"`
	Registrasi        []MKodeRegistrasi      `gorm:"foreignKey:IdStatusKendaraan;references:IdStatusKendaraan"`
	RegistrasiKhusus  []MKodeRegistrasiKhusu `gorm:"foreignKey:IdStatusKendaraan;references:IdStatusKendaraan"`
	IdStatusAktif     int                    `gorm:"default:1" json:"id_status_aktif"`
	GormModel
}
