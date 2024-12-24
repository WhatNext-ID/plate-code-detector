package controllers

import (
	"context"
	"net/http"
	"server/database"
	platecode "server/models/plate-code"
	"server/utils"

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
	vehicleCategory.ColorCriteria = utils.StringArray(newCategory.ColorCriteria)

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
