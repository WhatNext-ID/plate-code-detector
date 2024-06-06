package controllers

import (
	"net/http"
	"server/database"
	"server/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostKodeWilayah(ctx *gin.Context) {
	db := database.GetDB()
	KodeWilayah := models.MKodeWilayah{}

	ctx.ShouldBindJSON(&KodeWilayah)

	newId := uuid.New()

	KodeWilayah.IdKodeWilayah = newId

	if err := db.Create(&KodeWilayah).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":             "Kode Wilayah Berhasil Ditambahkan",
		"kodeWilayah":         KodeWilayah.KodeWilayah,
		"letakKodeRegistrasi": KodeWilayah.LetakKodeRegistrasi,
		"keterangan":          KodeWilayah.Keterangan,
	})
}
