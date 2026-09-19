package auth

import (
	"context"
	"errors"
	"net/http"
	"plate-server/database"
	"plate-server/helpers"
	"plate-server/models/user"
	errorhandling "plate-server/utils/error-handling"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbgorm"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func register(data Register) (string, error) {
	db := database.GetDB()

	userRegist := user.User{
		IdUser:    uuid.New(),
		UserName:  data.UserName,
		SecretKey: data.SecretKey,
	}

	userRole := user.UserRole{}

	err := crdbgorm.ExecuteTx(
		context.Background(),
		db,
		nil,
		func(tx *gorm.DB) error {
			if err := tx.
				Where("role_name = ?", data.UserRole).
				First(&userRole).Error; err != nil {

				if errors.Is(err, gorm.ErrRecordNotFound) {
					return &errorhandling.ServiceError{
						StatusCode: http.StatusBadRequest,
						ErrorName:  "Bad Request",
						Message:    "Role pengguna tidak ditemukan",
					}
				}

				return err
			}

			userRegist.IdRole = userRole.IdRole

			if err := tx.Create(&userRegist).Error; err != nil {
				return err
			}

			return nil
		},
	)

	if err != nil {
		return "", err
	}

	return "Pengguna berhasil ditambahkan", nil
}

func login(data Login) (string, error) {
	db := database.GetDB()

	userData := user.User{}

	err := db.
		Where("user_name = ?", data.UserName).
		Take(&userData).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", &errorhandling.ServiceError{
				StatusCode: http.StatusUnauthorized,
				ErrorName:  "Unauthorized",
				Message:    "Nama tidak ditemukan",
			}
		}

		return "", err
	}

	comparePass := helpers.ComparePass(
		[]byte(userData.SecretKey),
		[]byte(data.SecretKey),
	)

	if !comparePass {
		return "", &errorhandling.ServiceError{
			StatusCode: http.StatusUnauthorized,
			ErrorName:  "Unauthorized",
			Message:    "Kata sandi salah",
		}
	}

	token := helpers.GenerateToken(
		userData.IdUser,
		userData.UserName,
	)

	if token == "" {
		return "", &errorhandling.ServiceError{
			StatusCode: http.StatusInternalServerError,
			ErrorName:  "Internal Server Error",
			Message:    "Failed to generate token",
		}
	}

	return token, nil
}
