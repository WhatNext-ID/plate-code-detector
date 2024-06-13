package controllers

import (
	"net/http"
	"server/database"
	"server/models"

	"github.com/gin-gonic/gin"
)

type Code struct {
	KodeWilayah    string `json:"kode_wilayah"`
	KodeRegistrasi string `json:"kode_registrasi"`
}

func CheckCode(ctx *gin.Context) {
	db := database.GetDB()
	CodeData := Code{}
	var (
		KodeWilayah models.MKodeWilayah
		//KodeRegistrasi     models.MKodeRegistrasi
		ListKodeRegistrasi []models.MKodeRegistrasi
	)

	ctx.ShouldBindJSON(&CodeData)

	if err := db.First(&KodeWilayah, "kode_wilayah = ?", &CodeData.KodeWilayah).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := db.Find(&ListKodeRegistrasi, "id_kode_wilayah = ?", &KodeWilayah.IdKodeWilayah).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"data": ListKodeRegistrasi,
	})
}
