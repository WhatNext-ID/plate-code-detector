package controllers

import (
	"context"
	"net/http"
	"server/database"
	"server/models"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbgorm"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

func UpdateKodeWilayah(ctx *gin.Context) {
	db := database.GetDB()
	Admin := ctx.MustGet("userAdmin").(jwt.MapClaims)
	var (
		KodeWilayah       models.MKodeWilayah
		UpdateKodeWilayah models.MKodeWilayah
		User              models.MAdmin
	)
	id := ctx.Param("id")

	ctx.ShouldBindJSON((&UpdateKodeWilayah))

	if err := crdbgorm.ExecuteTx(context.Background(), db, nil, func(tx *gorm.DB) error {
		if err := tx.Where("id_admin = ?", Admin["id"]).First(&User).Error; err != nil {
			return err
		}

		if err := tx.Where("id_kode_wilayah = ?", id).First(&KodeWilayah).Error; err != nil {
			return err
		}

		if err := tx.Where("id_kode_wilayah = ?", KodeWilayah.IdKodeWilayah).Updates(&UpdateKodeWilayah).Error; err != nil {
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

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data status kendaraan berhasil diperbarui",
	})
}

func GetAllKodeWilayah(ctx *gin.Context) {
	db := database.GetDB()
	var ListKodeWilayah []models.MKodeWilayah

	if err := db.Find(&ListKodeWilayah).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"data": ListKodeWilayah,
	})
}

func GetKodeWilayahById(ctx *gin.Context) {
	db := database.GetDB()
	KodeWilayah := models.MKodeWilayah{}
	id := ctx.Param("id")

	if err := db.Where("id_kdoe_wilayah = ?", id).First(&KodeWilayah).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"kodeWilayah":         KodeWilayah.KodeWilayah,
		"letakKodeRegistrasi": KodeWilayah.LetakKodeRegistrasi,
		"keterangan":          KodeWilayah.Keterangan,
	})
}
