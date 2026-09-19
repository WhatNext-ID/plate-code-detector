package checkplate

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckPlateData(ctx *gin.Context) {
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
	checkDetailData, err := checkDetailPlate(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	// Construct the final JSON response
	ctx.JSON(http.StatusOK, gin.H{
		"data": checkDetailData,
	})
}
