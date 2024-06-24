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

	if err := ctx.ShouldBindJSON(&KodeWilayah); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

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
		"message":     "Kode Wilayah Berhasil Ditambahkan",
		"kodeWilayah": KodeWilayah.KodeWilayah,
		"keterangan":  KodeWilayah.Keterangan,
	})
}

func UpdateKodeWilayah(ctx *gin.Context) {
	Admin := ctx.MustGet("userAdmin").(jwt.MapClaims)

	db := database.GetDB()
	var (
		KodeWilayah       models.MKodeWilayah
		UpdateKodeWilayah models.MKodeWilayah
		User              models.MAdmin
	)
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&UpdateKodeWilayah); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := crdbgorm.ExecuteTx(context.Background(), db, nil, func(tx *gorm.DB) error {
		if err := tx.First(&User, "id_admin = ?", Admin["id"]).Error; err != nil {
			return err
		}

		if err := tx.First(&KodeWilayah, "id_kode_wilayah = ?", id).Error; err != nil {
			return err
		}

		if err := tx.Model(&KodeWilayah).Where("id_kode_wilayah = ?", KodeWilayah.IdKodeWilayah).Updates(&UpdateKodeWilayah).Error; err != nil {
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
	var (
		KodeWilayah []models.MKodeWilayah
		result      []gin.H
	)

	if err := db.Find(&KodeWilayah).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	for _, kw := range KodeWilayah {
		result = append(result, gin.H{
			"id":         kw.IdKodeWilayah,
			"kode":       kw.KodeWilayah,
			"keterangan": kw.Keterangan,
		})
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"data": result,
	})
}

func GetKodeWilayahById(ctx *gin.Context) {
	db := database.GetDB()
	KodeWilayah := models.MKodeWilayah{}
	id := ctx.Param("id")

	if err := db.Where("id_kode_wilayah = ?", id).First(&KodeWilayah).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"kodeWilayah": KodeWilayah.KodeWilayah,
		"keterangan":  KodeWilayah.Keterangan,
	})
}
