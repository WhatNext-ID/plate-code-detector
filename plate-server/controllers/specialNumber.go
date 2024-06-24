package controllers

import (
	"net/http"
	"server/database"
	"server/models"

	"github.com/gin-gonic/gin"
)

func PostSpecialNumber(ctx *gin.Context) {
	db := database.GetDB()
	var (
		PostNumber  models.PostNomorKhusus
		NomorKhusus models.MNomorKhusu
	)

	if err := ctx.ShouldBindJSON(&PostNumber); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
}
