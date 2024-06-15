package controllers

import (
	"context"
	"net/http"
	"server/database"
	"server/models"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbgorm"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func PostSpecialRegisterCode(ctx *gin.Context) {
	db := database.GetDB()

	var (
		Post                 models.KodeRegisterKhususPost
		StatusKendaraan      models.MStatusKendaraan
		KodeRegistrasiKhusus models.MKodeRegistrasiKhusu
		KodeWilayah          models.MKodeWilayah
	)

	if err := ctx.ShouldBindJSON(&Post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	newId := uuid.New()
	KodeRegistrasiKhusus.IdKodeRegistrasi = newId
	KodeRegistrasiKhusus.KodeRegistrasi = Post.KodeRegistrasi

	if Post.WilayahHukum != nil {
		KodeRegistrasiKhusus.WilayahHukum = *Post.WilayahHukum
	}

	if err := crdbgorm.ExecuteTx(context.Background(), db, nil, func(tx *gorm.DB) error {
		if err := tx.First(&KodeWilayah, "kode_wilayah = ?", &Post.KodeWilayah).Error; err != nil {
			return err
		}

		if err := tx.First(&StatusKendaraan, "status_kendaraan = ?", &Post.StatusKendaraan).Error; err != nil {
			return err
		}

		KodeRegistrasiKhusus.IdStatusKendaraan = StatusKendaraan.IdStatusKendaraan
		KodeRegistrasiKhusus.IdKodeWilayah = KodeWilayah.IdKodeWilayah

		if err := tx.Model(&KodeRegistrasiKhusus).Create(&KodeRegistrasiKhusus).Error; err != nil {
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

	ctx.JSON(http.StatusCreated, gin.H{
		"message":        "Kode Wilayah Berhasil Ditambahkan",
		"kodeRegistrasi": KodeRegistrasiKhusus.KodeRegistrasi,
		"keterangan":     KodeWilayah.Keterangan,
		"wilayahHukum":   KodeRegistrasiKhusus.WilayahHukum,
	})
}
