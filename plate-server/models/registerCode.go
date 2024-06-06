package models

import "github.com/google/uuid"

type MKodeRegistrasi struct {
	IdKodeRegistrasi uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id_kode_registrasi"`
	KodeRegistrasi   string    `gorm:"not null" valid:"required~Harap tambahkan kode plat bagian belakang" json:"kode_registrasi"`
	WilayahHukum     string    `gorm:"not null" valid:"required~Harap tambahkan wilayah hukum plat" json:"wilayah_hukum"`
	Keterangan       string    `gorm:"not null" valid:"required~Harap tambahkan keterangan kode plat bagian belakang" json:"keterangan"`
	IdKodeWilayah    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id_kode_wilayah"`
	IdStatusAktif    int       `gorm:"default:1" json:"id_status_aktif"`
	GormModel
}
