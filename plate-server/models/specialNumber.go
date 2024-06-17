package models

import "github.com/google/uuid"

type MNomorKhusu struct {
	IdNomorKhusus uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id_nomor_khusus"`
	GormModel
}
