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

type KodeRegisterPost struct {
	KodeRegistrasi  string  `json:"kode_registrasi"`
	WilayahHukum    *string `json:"wilayah_hukum"`
	Keterangan      string  `json:"keterangan"`
	KodeWilayah     string  `json:"kode_wilayah"`
	StatusKendaraan *string `json:"status_kendaraan"`
}

func PostKodeRegister(ctx *gin.Context) {
	db := database.GetDB()
	var (
		Post            KodeRegisterPost
		StatusKendaraan models.MStatusKendaraan
		KodeRegistrasi  models.MKodeRegistrasi
		KodeWilayah     models.MKodeWilayah
	)

	if err := ctx.ShouldBindJSON(&Post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	newId := uuid.New()
	KodeRegistrasi.IdKodeRegistrasi = newId
	KodeRegistrasi.KodeRegistrasi = Post.KodeRegistrasi
	KodeRegistrasi.Keterangan = Post.Keterangan

	if Post.WilayahHukum != nil {
		KodeRegistrasi.WilayahHukum = *Post.WilayahHukum
	}

	if err := crdbgorm.ExecuteTx(context.Background(), db, nil, func(tx *gorm.DB) error {
		if err := tx.First(&KodeWilayah, "kode_wilayah = ?", &Post.KodeWilayah).Error; err != nil {
			return err
		}

		if err := tx.First(&StatusKendaraan, "status_kendaraan = ?", &Post.StatusKendaraan).Error; err != nil {
			return err
		}

		if err := tx.First(&StatusKendaraan, "status_kendaraan = ?", &Post.StatusKendaraan).Error; err != nil {
			return err
		}

		KodeRegistrasi.IdStatusKendaraan = StatusKendaraan.IdStatusKendaraan
		KodeRegistrasi.IdKodeWilayah = KodeWilayah.IdKodeWilayah

		if err := tx.Create(&KodeRegistrasi).Error; err != nil {
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
		"kodeRegistrasi": KodeRegistrasi.KodeRegistrasi,
		"keterangan":     KodeRegistrasi.Keterangan,
		"wilayahHukum":   KodeRegistrasi.WilayahHukum,
	})
}
