package controllers

import (
	"net/http"
	"server/database"
	"server/models"

	"github.com/gin-gonic/gin"
)

func handleKendaraanDinas(ctx *gin.Context, CodeData models.Code) {
	db := database.GetDB()
	var result models.KendaraanKhususResult

	if err := db.Table("m_kode_registrasi_khusus").
		Select("m_kode_registrasi_khusus.kode_registrasi, m_kode_registrasi_khusus.wilayah_hukum, m_status_kendaraans.status_kendaraan, m_status_kendaraans.keterangan").
		Joins("JOIN m_status_kendaraans ON m_status_kendaraans.id_status_kendaraan = m_kode_registrasi_khusus.id_status_kendaraan").
		Where("m_kode_registrasi_khusus.kode_registrasi = ? AND m_status_kendaraans.status_kendaraan = ?", &CodeData.KodeRegistrasi, &CodeData.StatusKendaraan).
		First(&result).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"kodeKhusus": result,
	})
}

func handleKendaraanPribadi(ctx *gin.Context, CodeData models.Code) {
	db := database.GetDB()
	var (
		KodeWilayah models.MKodeWilayah
		result      models.KendaraanPribadiResult
	)

	if err := db.First(&KodeWilayah, "kode_wilayah = ?", CodeData.KodeWilayah).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := db.Table("m_kode_registrasis").
		Select("m_kode_registrasis.wilayah_hukum, m_kode_registrasis.kode_awal, m_kode_registrasis.kode_akhir, m_kode_registrasis.kode_alias, m_kode_wilayahs.keterangan as nama_provinsi").
		Joins("JOIN m_kode_wilayahs ON m_kode_registrasis.id_kode_wilayah = m_kode_wilayahs.id_kode_wilayah").
		Joins("JOIN m_status_kendaraans ON m_kode_registrasis.id_status_kendaraan = m_status_kendaraans.id_status_kendaraan").
		Where("m_kode_wilayahs.id_kode_wilayah = ? AND m_status_kendaraans.status_kendaraan = ? AND (m_kode_registrasis.kode_awal = ? OR m_kode_registrasis.kode_akhir = ? OR m_kode_registrasis.kode_alias = ? OR m_kode_registrasis.kode_alias = ?)",
			KodeWilayah.IdKodeWilayah, CodeData.StatusKendaraan, CodeData.KodeAwal, CodeData.KodeAkhir, CodeData.KodeAliasAwal, CodeData.KodeAliasAkhir).
		Take(&result).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	var kodeToUse string
	if result.KodeAwal == CodeData.KodeAwal || result.KodeAlias == CodeData.KodeAliasAwal {
		kodeToUse = CodeData.KodeAkhir
	} else if result.KodeAkhir == CodeData.KodeAkhir || result.KodeAlias == CodeData.KodeAliasAkhir {
		kodeToUse = CodeData.KodeAwal
	}

	var KodeKhusus models.KendaraanKhususResult
	if err := db.Table("m_kode_registrasi_khusus").
		Select("m_kode_registrasi_khusus.kode_registrasi, m_kode_registrasi_khusus.wilayah_hukum, m_status_kendaraans.status_kendaraan, m_status_kendaraans.keterangan").
		Joins("JOIN m_status_kendaraans ON m_status_kendaraans.id_status_kendaraan = m_kode_registrasi_khusus.id_status_kendaraan").
		Where("m_kode_registrasi_khusus.id_kode_wilayah = ? AND m_kode_registrasi_khusus.kode_registrasi = ?",
			KodeWilayah.IdKodeWilayah, kodeToUse).
		Take(&KodeKhusus).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"kodeRegistrasi": result,
		"kodeKhusus":     KodeKhusus,
	})
}

func CheckCode(ctx *gin.Context) {
	CodeData := models.Code{}

	if err := ctx.ShouldBindJSON(&CodeData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	switch CodeData.StatusKendaraan {
	case "Kendaraan Dinas":
		handleKendaraanDinas(ctx, CodeData)
	default:
		handleKendaraanPribadi(ctx, CodeData)
	}
}
