package controllers

import (
	"context"
	"net/http"
	"server/database"
	platecode "server/models/plate-code"
	"server/utils"
	"strings"
	"time"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbgorm"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VehicleType struct {
	VehicleType string `json:"vehicleType"`
}

type VehicleEngine struct {
	EngineType string `json:"engineType"`
}

type VehicleCategory struct {
	VehicleType   string   `json:"vehicleType"`
	VehicleEngine string   `json:"vehicleEngine"`
	ColorCriteria []string `json:"colorCriteria"`
}

type VehicleData struct {
	IdVehicleCategory uuid.UUID                 `json:"idVehicleCat"`
	VehicleType       string                    `json:"vehicleType"`
	VehicleEngine     string                    `json:"vehicleEngine"`
	ColorCriteria     utils.StringArrayResponse `json:"vehicleColorCriteria"`
	CreatedAt         *time.Time                `json:"createdAt"`
}

func CreateVehicleType(ctx *gin.Context) {
	db := database.GetDB()
	vehicleType := platecode.VehicleType{}
	var newVehicle VehicleType

	if err := ctx.ShouldBindJSON(&newVehicle); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	newId := uuid.New()
	vehicleType.IdVehicleType = newId
	vehicleType.VehicleType = newVehicle.VehicleType

	if err := db.Create(&vehicleType).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Jenis kendaraan berhasil ditambahkan",
	})
}

func CreateVehicleEngine(ctx *gin.Context) {
	db := database.GetDB()
	vehicleEngine := platecode.VehicleEngine{}
	var newEngine VehicleEngine

	if err := ctx.ShouldBindJSON(&newEngine); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	newId := uuid.New()
	vehicleEngine.IdVehicleEngine = newId
	vehicleEngine.VehicleEngineType = newEngine.EngineType

	if err := db.Create(&vehicleEngine).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Jenis mesin kendaraan berhasil ditambahkan",
	})
}

func CreateVehicleCategory(ctx *gin.Context) {
	db := database.GetDB()
	vehicleCategory := platecode.VehicleCategory{}
	vehicleEngine := platecode.VehicleEngine{}
	vehicleType := platecode.VehicleType{}
	var newCategory VehicleCategory

	if err := ctx.ShouldBindJSON(&newCategory); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	newId := uuid.New()
	vehicleCategory.IdVehicleCategory = newId
	vehicleCategory.ColorCriteria = utils.StringArrayDB(newCategory.ColorCriteria)

	if err := crdbgorm.ExecuteTx(context.Background(), db, nil, func(tx *gorm.DB) error {
		if err := tx.Model(&vehicleType).First(&vehicleType, "vehicle_type = ?", &newCategory.VehicleType).Error; err != nil {
			return err
		}

		vehicleCategory.IdVehicleType = vehicleType.IdVehicleType

		if err := tx.Model(&vehicleEngine).First(&vehicleEngine, "vehicle_engine_type = ?", &newCategory.VehicleEngine).Error; err != nil {
			return err
		}

		vehicleCategory.IdVehicleEngine = vehicleEngine.IdVehicleEngine

		if err := tx.Create(&vehicleCategory).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Kategori kendaraan berhasil ditambahkan",
	})
}

func GetVehicle(ctx *gin.Context) {
	db := database.GetDB()
	var vehicles []VehicleData
	var result = make([]gin.H, 0)

	// Query with ordering
	if err := db.Table("vehicle_categories AS vehicle").
		Select("vehicle.id_vehicle_category, types.vehicle_type, engines.vehicle_engine_type AS vehicle_engine, vehicle.color_criteria, vehicle.created_at").
		Joins("JOIN vehicle_types AS types ON types.id_vehicle_type = vehicle.id_vehicle_type").
		Joins("JOIN vehicle_engines AS engines ON engines.id_vehicle_engine = vehicle.id_vehicle_engine").
		Where("vehicle.id_status = ?", 1).
		Find(&vehicles).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	// Process the result
	for _, vehicle := range vehicles {
		// Extract and clean up color_criteria
		colorCriteria := ""
		if len(vehicle.ColorCriteria) > 0 {
			// Remove { } and replace , with ", "
			colorCriteria = strings.Trim(vehicle.ColorCriteria[0], "{}")
			colorCriteria = strings.ReplaceAll(colorCriteria, ",", ", ")
		}

		result = append(result, gin.H{
			"idVehicleCat":         vehicle.IdVehicleCategory,
			"vehicleType":          vehicle.VehicleType,
			"vehicleEngine":        vehicle.VehicleEngine,
			"vehicleColorCriteria": colorCriteria, // Store as cleaned string
			"createdAt":            vehicle.CreatedAt,
		})
	}

	// Return the JSON response
	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
