package regionplate

import (
	"errors"
	"net/http"
	"plate-server/database"
	platecode "plate-server/models/plate-code"
	errorhandling "plate-server/utils/error-handling"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func createRegionCode(data RegionCode) (string, error) {
	db := database.GetDB()

	var existingRegion platecode.RegionPlateCode

	err := db.
		Where(
			"region_code = ? AND region_area = ?",
			data.Code,
			data.Area,
		).
		First(&existingRegion).
		Error

	if err == nil {
		return "", &errorhandling.ServiceError{
			StatusCode: http.StatusConflict,
			ErrorName:  "Conflict",
			Message:    "Data region sudah tersedia",
		}
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", err
	}

	regionCode := platecode.RegionPlateCode{
		IdRegionCode: uuid.New(),
		RegionCode:   data.Code,
		RegionArea:   data.Area,
		Note:         data.Note,
	}

	if err := db.Create(&regionCode).Error; err != nil {
		return "", err
	}

	return "New region successfully created", nil
}

func getRegionCode() ([]RegionCodeResponse, error) {
	db := database.GetDB()

	var regionCodes []platecode.RegionPlateCode

	err := db.
		Select(`
			region_code,
			region_area,
			note,
			created_at
		`).
		Order("created_at DESC").
		Where("id_status = ?", 1).
		Find(&regionCodes).
		Error

	if err != nil {
		return nil, err
	}

	result := make([]RegionCodeResponse, 0, len(regionCodes))

	for _, rc := range regionCodes {
		result = append(result, RegionCodeResponse{
			RegionCode:  rc.RegionCode,
			RegionArea:  rc.RegionArea,
			RegionNote:  rc.Note,
			RegionAdded: rc.CreatedAt,
		})
	}

	return result, nil
}
