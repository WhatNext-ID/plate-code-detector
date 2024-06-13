package models

import (
	"fmt"
	"server/helpers"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MAdmin struct {
	IdAdmin       uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id_admin"`
	NamaAdmin     string    `gorm:"not null" valid:"required~Harap masukkan nama admin anda, stringlength(8|20)~Nama admin terdiri dari 8 sampai 20 karakter" json:"nama_admin"`
	SandiAdmin    string    `gorm:"not null" valid:"required~Harap masukkan kata sandi admin anda" json:"pw_admin"`
	IdStatusAktif int       `gorm:"default:1" json:"id_status_aktif"`
	GormModel
}

func (a *MAdmin) BeforeCreate(tx *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(a); err != nil {
		return fmt.Errorf("%w", err)
	}

	a.SandiAdmin = helpers.HashPass(a.SandiAdmin)
	fmt.Println(a.SandiAdmin)
	return nil
}
