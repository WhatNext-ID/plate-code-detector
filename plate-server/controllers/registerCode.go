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
	KodeAwal        *string `json:"kode_awal"`
	KodeAkhir       *string `json:"kode_akhir"`
	KodeAlias       *string `json:"kode_alias"`
	WilayahHukum    *string `json:"wilayah_hukum"`
	Keterangan      string  `json:"keterangan"`
	KodeWilayah     string  `json:"kode_wilayah"`
	StatusKendaraan string  `json:"status_kendaraan"`
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

	if Post.KodeAwal != nil {
		KodeRegistrasi.KodeAwal = *Post.KodeAwal
	}

	if Post.KodeAkhir != nil {
		KodeRegistrasi.KodeAkhir = *Post.KodeAkhir
	}

	if Post.KodeAlias != nil {
		KodeRegistrasi.KodeAlias = *Post.KodeAlias
	}

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

		KodeRegistrasi.IdStatusKendaraan = StatusKendaraan.IdStatusKendaraan
		KodeRegistrasi.IdKodeWilayah = KodeWilayah.IdKodeWilayah

		if err := tx.Model(&KodeRegistrasi).Create(&KodeRegistrasi).Error; err != nil {
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
		"message":      "Kode Wilayah Berhasil Ditambahkan",
		"kodeAwal":     KodeRegistrasi.KodeAwal,
		"kodeAkhir":    KodeRegistrasi.KodeAkhir,
		"kodeKhusus":   KodeRegistrasi.KodeAlias,
		"keterangan":   KodeWilayah.Keterangan,
		"wilayahHukum": KodeRegistrasi.WilayahHukum,
	})
}
