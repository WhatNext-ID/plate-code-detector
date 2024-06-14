package controllers

import (
	"net/http"
	"server/database"
	"server/models"

	"github.com/gin-gonic/gin"
)

type Code struct {
	KodeWilayah     string `json:"kode_wilayah"`
	KodeAwal        string `json:"kode_awal"`
	KodeAkhir       string `json:"kode_akhir"`
	KodeAliasAwal   string `json:"kode_alias_awal"`
	KodeAliasAkhir  string `json:"kode_alias_akhir"`
	KodeRegistrasi  string `json:"kode_registrasi"`
	StatusKendaraan string `json:"status_kendaraan"`
}

func NonPrivateVehicle(ctx *gin.Context) {
	db := database.GetDB()
	KodeKhusus := models.MKodeRegistrasiKhusu{}

}

func CheckCode(ctx *gin.Context) {
	db := database.GetDB()
	CodeData := Code{}
	var (
		KodeWilayah    models.MKodeWilayah
		KodeKhusus     models.MKodeRegistrasiKhusu
		KodeRegistrasi models.MKodeRegistrasi
	)

	if err := ctx.ShouldBindJSON(&CodeData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := db.First(&KodeWilayah, "kode_wilayah = ?", CodeData.KodeWilayah).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := db.Where(
		"id_kode_wilayah = ? AND (kode_awal = ? OR kode_akhir = ? OR kode_alias = ? OR kode_alias = ?)",
		KodeWilayah.IdKodeWilayah, CodeData.KodeAwal, CodeData.KodeAkhir, CodeData.KodeAliasAwal, CodeData.KodeAliasAkhir,
	).First(&KodeRegistrasi).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := db.Where(
		"id_kode_wilayah = ? AND (kode_registrasi = ? OR kode_registrasi = ?)",
		KodeWilayah.IdKodeWilayah, CodeData.KodeAwal, CodeData.KodeAkhir,
	).First(&KodeKhusus).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"kodeRegistrasi": KodeRegistrasi,
		"kodeKhusus":     KodeKhusus,
	})
}
