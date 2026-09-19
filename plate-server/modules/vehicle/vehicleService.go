package vehicle

import (
	"context"
	"errors"
	"net/http"
	"plate-server/database"
	platecode "plate-server/models/plate-code"
	"plate-server/utils"
	errorhandling "plate-server/utils/error-handling"
	"strings"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbgorm"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func createVehicleType(data VehicleType) (string, error) {
	db := database.GetDB()

	vehicleType := platecode.VehicleType{
		IdVehicleType: uuid.New(),
		VehicleType:   data.VehicleType,
	}

	if err := db.Create(&vehicleType).Error; err != nil {
		return "", err
	}

	return "Jenis kendaraan berhasil ditambahkan", nil
}

func createVehicleEngine(data VehicleEngine) (string, error) {
	db := database.GetDB()

	vehicleEngine := platecode.VehicleEngine{
		IdVehicleEngine:   uuid.New(),
		VehicleEngineType: data.EngineType,
	}

	if err := db.Create(&vehicleEngine).Error; err != nil {
		return "", err
	}

	return "Jenis mesin kendaraan berhasil ditambahkan", nil
}

func createVehicleCategory(data VehicleCategory) (string, error) {
	db := database.GetDB()

	vehicleCategory := platecode.VehicleCategory{
		IdVehicleCategory: uuid.New(),
		ColorCriteria:     utils.StringArrayDB(data.ColorCriteria),
	}

	err := crdbgorm.ExecuteTx(
		context.Background(),
		db,
		nil,
		func(tx *gorm.DB) error {
			var vehicleType platecode.VehicleType

			if err := tx.
				Where("vehicle_type = ?", data.VehicleType).
				First(&vehicleType).
				Error; err != nil {

				if errors.Is(err, gorm.ErrRecordNotFound) {
					return &errorhandling.ServiceError{
						StatusCode: http.StatusNotFound,
						ErrorName:  "Not Found",
						Message:    "Jenis kendaraan tidak ditemukan",
					}
				}

				return err
			}

			vehicleCategory.IdVehicleType = vehicleType.IdVehicleType

			var vehicleEngine platecode.VehicleEngine

			if err := tx.
				Where("vehicle_engine_type = ?", data.VehicleEngine).
				First(&vehicleEngine).
				Error; err != nil {

				if errors.Is(err, gorm.ErrRecordNotFound) {
					return &errorhandling.ServiceError{
						StatusCode: http.StatusNotFound,
						ErrorName:  "Not Found",
						Message:    "Jenis mesin kendaraan tidak ditemukan",
					}
				}

				return err
			}

			vehicleCategory.IdVehicleEngine = vehicleEngine.IdVehicleEngine

			if err := tx.Create(&vehicleCategory).Error; err != nil {
				return err
			}

			return nil
		},
	)

	if err != nil {
		return "", err
	}

	return "Kategori kendaraan berhasil ditambahkan", nil
}

func getVehicle() ([]VehicleResponse, error) {
	db := database.GetDB()

	var vehicles []VehicleData

	err := db.
		Table("vehicle_categories AS vehicle").
		Select(`
			types.vehicle_type,
			engines.vehicle_engine_type AS vehicle_engine,
			vehicle.color_criteria,
			vehicle.created_at
		`).
		Joins(`
			JOIN vehicle_types AS types
			ON types.id_vehicle_type = vehicle.id_vehicle_type
		`).
		Joins(`
			JOIN vehicle_engines AS engines
			ON engines.id_vehicle_engine = vehicle.id_vehicle_engine
		`).
		Where("vehicle.id_status = ?", 1).
		Find(&vehicles).
		Error

	if err != nil {
		return nil, err
	}

	result := make([]VehicleResponse, 0, len(vehicles))

	for _, vehicle := range vehicles {
		colorCriteria := ""

		if len(vehicle.ColorCriteria) > 0 {
			colorCriteria = strings.Trim(vehicle.ColorCriteria[0], "{}")
			colorCriteria = strings.ReplaceAll(
				colorCriteria,
				",",
				", ",
			)
		}

		result = append(result, VehicleResponse{
			VehicleType:          vehicle.VehicleType,
			VehicleEngine:        vehicle.VehicleEngine,
			VehicleColorCriteria: colorCriteria,
			CreatedAt:            vehicle.CreatedAt,
		})
	}

	return result, nil
}
