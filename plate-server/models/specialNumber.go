package models

import "github.com/google/uuid"

type MNomorKhusu struct {
	IdNomorKhusus     uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id_nomor_khusus"`
	NomorAwal         uint      `gorm:"null;default:null" json:"nomor_awal"`
	NomorAkhir        uint      `gorm:"null;default:null" json:"nomor_akhir"`
	IdWilayah         uuid.UUID `gorm:"not null" json:"id_wilayah_pengguna"`
	IdStatusAktif     int       `gorm:"default:1" json:"id_status_aktif"`
	IdStatusKendaraan uuid.UUID `gorm:"not null;type:uuid;default:uuid_generate_v4()" json:"id_status_kendaraan"`
	GormModel
}

type PostNomorKhusus struct {
	NomorAwal       uint   `json:"nomor_awal"`
	NomorAkhir      uint   `json:"nomor_akhir"`
	KodeWilayah     string `json:"kode_wilayah"`
	KodeRegistrasi  string `json:"kode_registrasi"`
	StatusKendaraan string `json:"status_kendaraan"`
}
