package controllers

import (
	"database/sql"
	"net/http"
	"plate-server/database"
	platecode "plate-server/models/plate-code"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type DataCode struct {
	BasePlateColor       string  `json:"basePlateColor"`
	TextPlateColor       string  `json:"textPlateColor"`
	AdditionalPlateColor *string `json:"additionalPlateColor"`
	RegionCode           string  `json:"regionCode"`
	RegisterFirstCode    string  `json:"registerFirstCode"`
	RegisterLastCode     string  `json:"registerLastCode"`
	RegisterCode         string  `json:"registerCode"`
}

// Check Vehicle Region
func CheckVehicleRegion(db *gorm.DB, region DataCode) (map[string]interface{}, uuid.UUID, error) {
	var vehicleRegion platecode.RegionPlateCode

	err := db.Table("region_plate_codes").Select("id_region_code, region_code, region_area, note").
		Where("region_code = ? AND id_status = ?", region.RegionCode, 1).
		First(&vehicleRegion).Error

	if err != nil {
		return nil, uuid.UUID{}, err
	}

	if err == gorm.ErrRecordNotFound {
		return nil, uuid.UUID{}, nil
	}

	// Return the result as a map object
	result := map[string]interface{}{
		"regionCode": vehicleRegion.RegionCode,
		"regionArea": vehicleRegion.RegionArea,
		"note":       vehicleRegion.Note,
	}

	return result, vehicleRegion.IdRegionCode, nil
}

// CheckVehicleRegister searches for a vehicle register based on priority order
func CheckVehicleRegister(db *gorm.DB, idRegion uuid.UUID, register DataCode) (map[string]interface{}, error) {
	var (
		registerCodePosition []sql.NullInt64
		registerCode         platecode.RegisterPlateCode
	)

	// Fetch distinct code positions
	err := db.Table("register_plate_codes").
		Select("DISTINCT code_position").
		Where("id_region_code = ?", idRegion).
		Pluck("code_position", &registerCodePosition).Error

	if err != nil {
		return nil, err
	}

	if len(registerCodePosition) == 0 {
		return nil, nil // No records found
	}

	// Convert sql.NullInt64 to []*int to safely handle NULL values
	var codePositions []*int
	for _, cp := range registerCodePosition {
		if cp.Valid {
			val := int(cp.Int64)
			codePositions = append(codePositions, &val)
		} else {
			codePositions = append(codePositions, nil)
		}
	}

	// Loop through all code positions
	for _, codePosition := range codePositions {
		if codePosition == nil {
			// If codePosition is null, use registerCode
			err = db.Table("register_plate_codes").
				Select("register_code, register_city, note").
				Where("id_region_code = ? AND register_code = ?", idRegion, register.RegisterCode).
				First(&registerCode).Error

			if err == nil {
				return map[string]interface{}{
					"register_code": registerCode.RegisterCode,
					"register_city": registerCode.RegisterCity,
					"note":          registerCode.Note,
				}, nil
			}
		}

		if codePosition != nil && *codePosition == 0 {
			// If codePosition is 0, use registerFirstCode
			err = db.Table("register_plate_codes").
				Select("register_code, register_city, note").
				Where("id_region_code = ? AND code_position = ? AND register_code = ?", idRegion, *codePosition, register.RegisterFirstCode).
				First(&registerCode).Error

			if err == nil {
				return map[string]interface{}{
					"register_code": registerCode.RegisterCode,
					"register_city": registerCode.RegisterCity,
					"note":          registerCode.Note,
				}, nil
			}
		}

		if codePosition != nil && (*codePosition == 1 || *codePosition == 2) {
			// If codePosition is 1 or 2, use registerLastCode
			err = db.Table("register_plate_codes").
				Select("register_code, register_city, note").
				Where("id_region_code = ? AND code_position = ? AND register_code = ?", idRegion, *codePosition, register.RegisterLastCode).
				First(&registerCode).Error

			if err == nil {
				return map[string]interface{}{
					"register_code": registerCode.RegisterCode,
					"register_city": registerCode.RegisterCity,
					"note":          registerCode.Note,
				}, nil
			}
		}
	}

	return nil, nil
}

// Check Vehicle Type or Status by Plate Color
func CheckVehicleStatus(db *gorm.DB, status DataCode) (map[string]interface{}, error) {
	var vehicleStatus struct {
		VehicleType   *string `json:"vehicleType"`
		VehicleEngine *string `json:"vehicleEngine"`
	}

	// Ensure AdditionalPlateColor is only included if not nil
	colorCriteria := []string{status.BasePlateColor, status.TextPlateColor}
	if status.AdditionalPlateColor != nil {
		colorCriteria = append(colorCriteria, *status.AdditionalPlateColor)
	}

	// Query Database
	err := db.Table("vehicle_categories AS vehicle").
		Select("types.vehicle_type, engines.vehicle_engine_type AS vehicle_engine").
		Joins("JOIN vehicle_types AS types ON types.id_vehicle_type = vehicle.id_vehicle_type").
		Joins("JOIN vehicle_engines AS engines ON engines.id_vehicle_engine = vehicle.id_vehicle_engine").
		Where("vehicle.color_criteria @> ? AND vehicle.id_status = ?", pq.Array(colorCriteria), 1).
		Find(&vehicleStatus).Error

	if err != nil {
		return nil, err
	}

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	// Return the result as a map object
	result := map[string]interface{}{
		"vehicleType":   vehicleStatus.VehicleType,
		"vehicleEngine": vehicleStatus.VehicleEngine,
	}

	return result, nil
}

// Check Plate Data
func CheckPlateData(ctx *gin.Context) {
	db := database.GetDB()
	body := DataCode{}

	// Bind JSON Request Body
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	// Get Vehicle Status Data
	vehicleStatus, err := CheckVehicleStatus(db, body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	// Get Vehicle Region Data
	vehicleRegion, idRegion, err := CheckVehicleRegion(db, body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	// Get Vehicle Register Area
	vehicleRegister, err := CheckVehicleRegister(db, idRegion, body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	// Construct the final JSON response
	ctx.JSON(http.StatusOK, gin.H{
		"status":   vehicleStatus,
		"region":   vehicleRegion,
		"register": vehicleRegister,
	})
}
