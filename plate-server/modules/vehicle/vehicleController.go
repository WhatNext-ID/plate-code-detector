package vehicle

import (
	"net/http"
	errorhandling "plate-server/utils/error-handling"

	"github.com/gin-gonic/gin"
)

func CreateVehicleType(ctx *gin.Context) {
	var newVehicle VehicleType

	if err := ctx.ShouldBindJSON(&newVehicle); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	result, err := createVehicleType(newVehicle)
	if err != nil {
		errorhandling.HandleServiceError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": result,
	})
}

func CreateVehicleEngine(ctx *gin.Context) {
	var newEngine VehicleEngine

	if err := ctx.ShouldBindJSON(&newEngine); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	result, err := createVehicleEngine(newEngine)
	if err != nil {
		errorhandling.HandleServiceError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": result,
	})
}

func CreateVehicleCategory(ctx *gin.Context) {
	var newCategory VehicleCategory

	if err := ctx.ShouldBindJSON(&newCategory); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	result, err := createVehicleCategory(newCategory)
	if err != nil {
		errorhandling.HandleServiceError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": result,
	})
}

func GetVehicle(ctx *gin.Context) {
	result, err := getVehicle()
	if err != nil {
		errorhandling.HandleServiceError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
