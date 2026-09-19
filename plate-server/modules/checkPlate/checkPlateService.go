package checkplate

import (
	"database/sql"
	"errors"
	"net/http"
	"plate-server/database"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// Check Vehicle Region
func checkVehicleRegion(db *gorm.DB, region string) (*VehicleRegionResponse, uuid.UUID, error) {
	vehicleRegionDb := VehicleRegionQuery{}

	err := db.Table("region_plate_codes").Select("id_region_code, region_code, region_area, note").
		Where("region_code = ? AND id_status = ?", region, 1).
		First(&vehicleRegionDb).Error

	if err != nil {
		return nil, uuid.UUID{}, err
	}

	if err == gorm.ErrRecordNotFound {
		return nil, uuid.UUID{}, nil
	}

	vehicleRegion := &VehicleRegionResponse{
		RegionCode: vehicleRegionDb.RegionCode,
		RegionArea: vehicleRegionDb.RegionArea,
		Note:       vehicleRegionDb.Note,
	}

	return vehicleRegion, vehicleRegionDb.IdRegionCode, nil
}

// CheckVehicleRegister searches for a vehicle register based on priority order
func checkVehicleRegister(db *gorm.DB, idRegion uuid.UUID, register VehicleRegisterParam) (*VehicleRegister, error) {
	var registerCodePosition []sql.NullInt64

	vehicleRegister := VehicleRegister{}

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
				First(&vehicleRegister).Error

			if err == nil {
				return &vehicleRegister, nil
			}
		}

		if codePosition != nil && *codePosition == 0 {
			// If codePosition is 0, use registerFirstCode
			err = db.Table("register_plate_codes").
				Select("register_code, register_city, note").
				Where("id_region_code = ? AND code_position = ? AND register_code = ?", idRegion, *codePosition, register.RegisterFirstCode).
				First(&vehicleRegister).Error

			if err == nil {
				return &vehicleRegister, nil
			}
		}

		if codePosition != nil && (*codePosition == 1 || *codePosition == 2) {
			// If codePosition is 1 or 2, use registerLastCode
			err = db.Table("register_plate_codes").
				Select("register_code, register_city, note").
				Where("id_region_code = ? AND code_position = ? AND register_code = ?", idRegion, *codePosition, register.RegisterLastCode).
				First(&vehicleRegister).Error

			if err == nil {
				return &vehicleRegister, nil
			}
		}
	}

	return nil, nil
}

// Check Vehicle Type or Status by Plate Color
func checkVehicleStatus(db *gorm.DB, status DataVehicleStatus) (*VehicleStatus, error) {
	vehicleStatus := VehicleStatus{}

	colorCriteria := []string{
		status.BasePlateColor,
		status.TextPlateColor,
	}

	if status.AdditionalPlateColor != nil {
		colorCriteria = append(colorCriteria, *status.AdditionalPlateColor)
	}

	err := db.Table("vehicle_categories AS vehicle").
		Select("types.vehicle_type, engines.vehicle_engine_type AS vehicle_engine").
		Joins("JOIN vehicle_types AS types ON types.id_vehicle_type = vehicle.id_vehicle_type").
		Joins("JOIN vehicle_engines AS engines ON engines.id_vehicle_engine = vehicle.id_vehicle_engine").
		Where("vehicle.color_criteria @> ? AND vehicle.id_status = ?", pq.Array(colorCriteria), 1).
		Find(&vehicleStatus).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &vehicleStatus, nil
}

// Check Plate Data
func checkDetailPlate(ctx *gin.Context, body DataCode) (map[string]interface{}, error) {
	db := database.GetDB()

	vehicleStatusInput := DataVehicleStatus{
		BasePlateColor:       body.BasePlateColor,
		AdditionalPlateColor: body.AdditionalPlateColor,
		TextPlateColor:       body.TextPlateColor,
	}
	vehicleRegisterParam := VehicleRegisterParam{
		RegisterCode:      body.RegisterCode,
		RegisterFirstCode: body.RegisterFirstCode,
		RegisterLastCode:  body.RegisterLastCode,
	}

	// Get Vehicle Status Data
	vehicleStatus, err := checkVehicleStatus(db, vehicleStatusInput)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return nil, err
	}

	// Get Vehicle Region Data
	vehicleRegion, idRegion, err := checkVehicleRegion(db, body.RegionCode)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return nil, err
	}

	// Get Vehicle Register Area
	vehicleRegister, err := checkVehicleRegister(db, idRegion, vehicleRegisterParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return nil, err
	}

	// Construct the final JSON response
	result := map[string]interface{}{
		"status":   vehicleStatus,
		"region":   vehicleRegion,
		"register": vehicleRegister,
	}

	return result, nil
}
