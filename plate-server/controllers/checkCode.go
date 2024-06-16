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

	query := db.Table("m_kode_registrasi_khusus").
		Select("m_kode_registrasi_khusus.kode_registrasi, m_kode_registrasi_khusus.wilayah_hukum, m_status_kendaraans.status_kendaraan, m_status_kendaraans.keterangan, m_kode_wilayahs.keterangan as nama_provinsi").
		Joins("JOIN m_status_kendaraans ON m_status_kendaraans.id_status_kendaraan = m_kode_registrasi_khusus.id_status_kendaraan").
		Joins("JOIN m_kode_wilayahs ON m_kode_wilayahs.id_kode_wilayah = m_kode_registrasi_khusus.id_kode_wilayah").
		Where("m_status_kendaraans.status_kendaraan = ?", CodeData.StatusKendaraan)

	if len(CodeData.KodeRegistrasi) != 1 {
		query = query.Where("m_kode_registrasi_khusus.kode_registrasi = ?", CodeData.KodeRegistrasi)
	} else {
		query = query.Where("m_kode_registrasi_khusus.kode_registrasi = ? OR m_kode_registrasi_khusus.kode_registrasi = ?", CodeData.KodeAwal, CodeData.KodeAkhir)
	}

	if CodeData.WilayahHukum != nil {
		query = query.Where("m_kode_registrasi_khusus.wilayah_hukum = ?", CodeData.WilayahHukum)
	}

	if err := query.First(&result).Error; err != nil {
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

func handleKendaraanKhusus(ctx *gin.Context, CodeData models.Code) {
	db := database.GetDB()
	var (
		KodeWilayah models.MKodeWilayah
		KodeKhusus  models.KendaraanKhususResult
		result      models.KendaraanPribadiResult
		kodeToUse   string
	)

	if err := db.First(&KodeWilayah, "kode_wilayah = ?", CodeData.KodeWilayah).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := db.Table("m_kode_registrasi_khusus").
		Select("m_kode_registrasi_khusus.kode_registrasi, m_kode_registrasi_khusus.wilayah_hukum, m_status_kendaraans.status_kendaraan, m_status_kendaraans.keterangan").
		Joins("JOIN m_status_kendaraans ON m_status_kendaraans.id_status_kendaraan = m_kode_registrasi_khusus.id_status_kendaraan").
		Where("m_kode_registrasi_khusus.id_kode_wilayah = ? AND (m_kode_registrasi_khusus.kode_registrasi = ? OR m_kode_registrasi_khusus.kode_registrasi = ? OR m_kode_registrasi_khusus.kode_registrasi = ?)",
			KodeWilayah.IdKodeWilayah, CodeData.KodeAwal, CodeData.KodeAkhir, CodeData.KodeRegistrasi).
		Take(&KodeKhusus).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if KodeKhusus.KodeRegistrasi == CodeData.KodeAwal {
		kodeToUse = CodeData.KodeAkhir
	} else if KodeKhusus.KodeRegistrasi == CodeData.KodeAkhir {
		kodeToUse = CodeData.KodeAwal
	}

	if err := db.Table("m_kode_registrasis").
		Select("m_kode_registrasis.wilayah_hukum, m_kode_registrasis.kode_awal, m_kode_registrasis.kode_akhir, m_kode_registrasis.kode_alias, m_kode_wilayahs.keterangan as nama_provinsi, m_status_kendaraans.status_kendaraan, m_status_kendaraans.keterangan").
		Joins("JOIN m_kode_wilayahs ON m_kode_registrasis.id_kode_wilayah = m_kode_wilayahs.id_kode_wilayah").
		Joins("JOIN m_status_kendaraans ON m_kode_registrasis.id_status_kendaraan = m_status_kendaraans.id_status_kendaraan").
		Where("m_kode_wilayahs.id_kode_wilayah = ? AND (m_kode_registrasis.kode_awal = ? OR m_kode_registrasis.kode_akhir = ? OR m_kode_registrasis.kode_alias = ? OR m_kode_registrasis.kode_alias = ?)",
			KodeWilayah.IdKodeWilayah, kodeToUse, kodeToUse, CodeData.KodeAliasAwal, CodeData.KodeAliasAkhir).
		Take(&result).Error; err != nil {
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
		Select("m_kode_registrasis.wilayah_hukum, m_kode_registrasis.kode_awal, m_kode_registrasis.kode_akhir, m_kode_registrasis.kode_alias, m_kode_wilayahs.keterangan as nama_provinsi, m_status_kendaraans.status_kendaraan, m_status_kendaraans.keterangan").
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

	ctx.JSON(http.StatusAccepted, gin.H{
		"kodeRegistrasi": result,
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
	case "Kendaraan Pribadi", "Free Trade Zone":
		handleKendaraanPribadi(ctx, CodeData)
	default:
		handleKendaraanKhusus(ctx, CodeData)
	}
}
