package registerplate

import (
	"net/http"
	errorhandling "plate-server/utils/error-handling"

	"github.com/gin-gonic/gin"
)

func CreateRegisterCode(ctx *gin.Context) {
	regionCode := ctx.Param("regionCode")

	var newRegister RegisterCode

	if err := ctx.ShouldBindJSON(&newRegister); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	result, err := createRegisterCode(
		regionCode,
		newRegister,
	)

	if err != nil {
		errorhandling.HandleServiceError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": result,
	})
}

func GetRegisterCode(ctx *gin.Context) {
	result, err := getRegisterCode()

	if err != nil {
		errorhandling.HandleServiceError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func GetRegisterCodeByRegionCode(ctx *gin.Context) {
	regionCode := ctx.Param("regionCode")

	result, err := getRegisterCodeByRegionCode(
		regionCode,
	)

	if err != nil {
		errorhandling.HandleServiceError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
