package checkplate

import (
	"database/sql"
	"errors"
	"net/http"
	"plate-server/database"
	errorhandling "plate-server/utils/error-handling"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// Check Vehicle Region
func checkVehicleRegion(
	db *gorm.DB,
	region string,
) (*VehicleRegionResponse, uuid.UUID, error) {

	var vehicleRegionDb VehicleRegionQuery

	err := db.Table("region_plate_codes").
		Select("id_region_code, region_code, region_area, note").
		Where("region_code = ? AND id_status = ?", region, 1).
		First(&vehicleRegionDb).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, uuid.Nil, &errorhandling.ServiceError{
			StatusCode: http.StatusNotFound,
			ErrorName:  "Not Found",
			Message:    "Kode wilayah kendaraan tidak ditemukan",
		}
	}

	if err != nil {
		return nil, uuid.Nil, err
	}

	vehicleRegion := &VehicleRegionResponse{
		RegionCode: vehicleRegionDb.RegionCode,
		RegionArea: vehicleRegionDb.RegionArea,
		Note:       vehicleRegionDb.Note,
	}

	return vehicleRegion, vehicleRegionDb.IdRegionCode, nil
}

// CheckVehicleRegister searches for a vehicle register based on priority order
func checkVehicleRegister(
	db *gorm.DB,
	idRegion uuid.UUID,
	register VehicleRegisterParam,
) (*VehicleRegister, error) {

	var registerCodePosition []sql.NullInt64
	var vehicleRegister VehicleRegister

	err := db.Table("register_plate_codes").
		Select("DISTINCT code_position").
		Where("id_region_code = ?", idRegion).
		Pluck("code_position", &registerCodePosition).
		Error

	if err != nil {
		return nil, err
	}

	if len(registerCodePosition) == 0 {
		return nil, &errorhandling.ServiceError{
			StatusCode: http.StatusNotFound,
			ErrorName:  "Not Found",
			Message:    "Kode registrasi kendaraan tidak ditemukan",
		}
	}

	for _, cp := range registerCodePosition {

		// code_position IS NULL
		if !cp.Valid {
			err = db.Table("register_plate_codes").
				Select("register_code, register_city, note").
				Where(`
					id_region_code = ?
					AND code_position IS NULL
					AND register_code = ?
				`, idRegion, register.RegisterCode).
				First(&vehicleRegister).
				Error

			if err == nil {
				return &vehicleRegister, nil
			}

			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, err
			}
		}

		// code_position = 0
		if cp.Valid && cp.Int64 == 0 {
			err = db.Table("register_plate_codes").
				Select("register_code, register_city, note").
				Where(`
					id_region_code = ?
					AND code_position = ?
					AND register_code = ?
				`, idRegion, 0, register.RegisterFirstCode).
				First(&vehicleRegister).
				Error

			if err == nil {
				return &vehicleRegister, nil
			}

			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, err
			}
		}

		// code_position = 1 or 2
		if cp.Valid && (cp.Int64 == 1 || cp.Int64 == 2) {
			err = db.Table("register_plate_codes").
				Select("register_code, register_city, note").
				Where(`
					id_region_code = ?
					AND code_position = ?
					AND register_code = ?
				`, idRegion, cp.Int64, register.RegisterLastCode).
				First(&vehicleRegister).
				Error

			if err == nil {
				return &vehicleRegister, nil
			}

			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, err
			}
		}
	}

	return nil, &errorhandling.ServiceError{
		StatusCode: http.StatusNotFound,
		ErrorName:  "Not Found",
		Message:    "Data registrasi kendaraan tidak ditemukan",
	}
}

// Check Vehicle Type or Status by Plate Color
func checkVehicleStatus(
	db *gorm.DB,
	status DataVehicleStatus,
) (*VehicleStatus, error) {

	var vehicleStatus VehicleStatus

	colorCriteria := []string{
		status.BasePlateColor,
		status.TextPlateColor,
	}

	if status.AdditionalPlateColor != nil {
		colorCriteria = append(
			colorCriteria,
			*status.AdditionalPlateColor,
		)
	}

	err := db.Table("vehicle_categories AS vehicle").
		Select(`
			types.vehicle_type AS vehicle_type,
			engines.vehicle_engine_type AS vehicle_engine
		`).
		Joins(`
			JOIN vehicle_types AS types
			ON types.id_vehicle_type = vehicle.id_vehicle_type
		`).
		Joins(`
			JOIN vehicle_engines AS engines
			ON engines.id_vehicle_engine = vehicle.id_vehicle_engine
		`).
		Where(
			"vehicle.color_criteria @> ? AND vehicle.id_status = ?",
			pq.Array(colorCriteria),
			1,
		).
		First(&vehicleStatus).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, &errorhandling.ServiceError{
			StatusCode: http.StatusNotFound,
			ErrorName:  "Not Found",
			Message:    "Jenis kendaraan tidak ditemukan",
		}
	}

	if err != nil {
		return nil, err
	}

	return &vehicleStatus, nil
}

// Check Plate Data
func checkDetailPlate(body DataCode) (*CheckDetail, error) {
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

	vehicleStatus, err := checkVehicleStatus(
		db,
		vehicleStatusInput,
	)
	if err != nil {
		return nil, err
	}

	vehicleRegion, idRegion, err := checkVehicleRegion(
		db,
		body.RegionCode,
	)
	if err != nil {
		return nil, err
	}

	vehicleRegister, err := checkVehicleRegister(
		db,
		idRegion,
		vehicleRegisterParam,
	)
	if err != nil {
		return nil, err
	}

	result := &CheckDetail{
		Region:   vehicleRegion,
		Register: vehicleRegister,
		Status:   vehicleStatus,
	}

	return result, nil
}
