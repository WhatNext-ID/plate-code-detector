package models

import "github.com/google/uuid"

type MKodeRegistrasi struct {
	IdKodeRegistrasi  uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id_kode_registrasi"`
	KodeAwal          string    `gorm:"null;default:null" valid:"required~Harap tambahkan kode registrasi" json:"kode_awal"`
	KodeAkhir         string    `gorm:"null;default:null" valid:"required~Harap tambahkan kode registrasi" json:"kode_akhir"`
	KodeAlias         string    `gorm:"null;default:null" valid:"required~Harap tambahkan kode registrasi" json:"kode_alias"`
	WilayahHukum      string    `gorm:"null;default:null" valid:"required~Harap tambahkan wilayah hukum plat" json:"wilayah_hukum"`
	IdKodeWilayah     uuid.UUID `gorm:"not null;type:uuid;default:uuid_generate_v4()" json:"id_kode_wilayah"`
	IdStatusKendaraan uuid.UUID `gorm:"not null;type:uuid;default:uuid_generate_v4()" json:"id_status_kendaraan"`
	IdStatusAktif     int       `gorm:"default:1" json:"id_status_aktif"`
	GormModel
}
