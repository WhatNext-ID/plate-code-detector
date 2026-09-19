package registerplate

import (
	"context"
	"errors"
	"net/http"
	"plate-server/database"
	platecode "plate-server/models/plate-code"
	errorhandling "plate-server/utils/error-handling"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbgorm"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func createRegisterCode(
	region string,
	data RegisterCode,
) (string, error) {
	db := database.GetDB()

	if region == "" {
		return "", &errorhandling.ServiceError{
			StatusCode: http.StatusBadRequest,
			ErrorName:  "Bad Request",
			Message:    "regionCode is required",
		}
	}

	var regionCode platecode.RegionPlateCode

	err := db.
		Where("region_code = ?", region).
		First(&regionCode).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", &errorhandling.ServiceError{
			StatusCode: http.StatusNotFound,
			ErrorName:  "Not Found",
			Message:    "Region code tidak ditemukan",
		}
	}

	if err != nil {
		return "", err
	}

	registerCode := platecode.RegisterPlateCode{
		IdRegisterCode: uuid.New(),
		IdRegionCode:   regionCode.IdRegionCode,
		RegisterCode:   data.RegisterCode,
		RegisterCity:   data.RegisterCity,
		Note:           data.Note,
		CodePosition:   data.CodePosition,
	}

	err = crdbgorm.ExecuteTx(
		context.Background(),
		db,
		nil,
		func(tx *gorm.DB) error {

			var existing platecode.RegisterPlateCode

			err := tx.
				Where(`
					register_code = ?
					AND register_city = ?
					AND id_region_code = ?
				`,
					data.RegisterCode,
					data.RegisterCity,
					regionCode.IdRegionCode,
				).
				First(&existing).
				Error

			if err == nil {
				return &errorhandling.ServiceError{
					StatusCode: http.StatusConflict,
					ErrorName:  "Conflict",
					Message:    "Data register sudah tersedia",
				}
			}

			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}

			if err := tx.Create(&registerCode).Error; err != nil {
				return err
			}

			return nil
		},
	)

	if err != nil {
		return "", err
	}

	return "New register successfully created", nil
}

func getRegisterCode() ([]RegisterCodeResponse, error) {
	db := database.GetDB()

	var registerCodes []platecode.RegisterPlateCode

	err := db.
		Select(`
			id_register_code,
			register_code,
			register_city,
			note,
			created_at
		`).
		Order("id_region_code ASC").
		Order("register_code ASC").
		Where("id_status = ?", 1).
		Find(&registerCodes).
		Error

	if err != nil {
		return nil, err
	}

	result := make(
		[]RegisterCodeResponse,
		0,
		len(registerCodes),
	)

	for _, rc := range registerCodes {
		result = append(result, RegisterCodeResponse{
			RegisterCode:  rc.RegisterCode,
			RegisterCity:  rc.RegisterCity,
			RegisterNote:  rc.Note,
			RegisterAdded: rc.CreatedAt,
		})
	}

	return result, nil
}

func getRegisterCodeByRegionCode(
	regionCode string,
) ([]RegisterCodeResponse, error) {

	if regionCode == "" {
		return nil, &errorhandling.ServiceError{
			StatusCode: http.StatusBadRequest,
			ErrorName:  "Bad Request",
			Message:    "regionCode is required",
		}
	}

	db := database.GetDB()

	var registerCodes []platecode.RegisterPlateCode

	err := db.
		Table("register_plate_codes AS register").
		Select(`
			register.id_register_code,
			register.register_code,
			register.register_city,
			register.note,
			register.created_at
		`).
		Joins(`
			JOIN region_plate_codes AS region
			ON region.id_region_code = register.id_region_code
		`).
		Where(`
			register.id_status = ?
			AND region.region_code = ?
		`, 1, regionCode).
		Order("register.register_code ASC").
		Find(&registerCodes).
		Error

	if err != nil {
		return nil, err
	}

	result := make(
		[]RegisterCodeResponse,
		0,
		len(registerCodes),
	)

	for _, rc := range registerCodes {
		result = append(result, RegisterCodeResponse{
			RegisterCode:  rc.RegisterCode,
			RegisterCity:  rc.RegisterCity,
			RegisterNote:  rc.Note,
			RegisterAdded: rc.CreatedAt,
		})
	}

	return result, nil
}
