package models

import "github.com/google/uuid"

type MKodeWilayah struct {
	IdKodeWilayah       uuid.UUID         `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id_kode_wilayah"`
	KodeWilayah         string            `gorm:"not null" valid:"required~Harap tambahkan kode plat bagian depan" json:"kode_wilayah"`
	Keterangan          string            `gorm:"not null" valid:"required~Harap tambahkan keterangan kode plat bagian depan" json:"keterangan"`
	LetakKodeRegistrasi int               `gorm:"not null" valid:"required~Harap tambahkan keterangan posisi kode plat bagian belakang" json:"letak_kode_registrasi"`
	IdStatusAktif       int               `gorm:"default:1" json:"id_status_aktif"`
	Registrasi          []MKodeRegistrasi `gorm:"foreignKey:IdKodeWilayah;references:IdKodeWilayah"`
	GormModel
}
