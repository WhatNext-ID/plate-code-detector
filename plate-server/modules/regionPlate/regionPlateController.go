package regionplate

import (
	"net/http"
	errorhandling "plate-server/utils/error-handling"

	"github.com/gin-gonic/gin"
)

func CreateRegionCode(ctx *gin.Context) {
	var newRegion RegionCode

	if err := ctx.ShouldBindJSON(&newRegion); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	result, err := createRegionCode(newRegion)
	if err != nil {
		errorhandling.HandleServiceError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": result,
	})
}

func GetRegionCode(ctx *gin.Context) {
	result, err := getRegionCode()
	if err != nil {
		errorhandling.HandleServiceError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
