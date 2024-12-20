package user

import (
	"fmt"
	"server/helpers"
	"server/models"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	IdUser    uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id_user"`
	UserName  string    `gorm:"unique;not null" valid:"required~Harap masukkan nama anda, stringlength(8|20)~Nama admin terdiri dari 8 sampai 20 karakter" json:"username"`
	SecretKey string    `gorm:"not null" valid:"required~Harap masukkan kata sandi anda" json:"secretkey"`
	IdRole    uuid.UUID `gorm:"type:uuid" json:"id_role"`
	models.GormModel
}

func (a *User) BeforeCreate(tx *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(a); err != nil {
		return fmt.Errorf("%w", err)
	}

	a.SecretKey = helpers.HashPass(a.SecretKey)
	fmt.Println(a)
	return nil
}
