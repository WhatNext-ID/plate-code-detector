package checkplate

import (
	"net/http"
	errorhandling "plate-server/utils/error-handling"

	"github.com/gin-gonic/gin"
)

func CheckPlateData(ctx *gin.Context) {
	var body DataCode

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	checkDetailData, err := checkDetailPlate(body)
	if err != nil {
		errorhandling.HandleServiceError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": checkDetailData,
	})
}

func CheckMultiplePlateData(ctx *gin.Context) {}
