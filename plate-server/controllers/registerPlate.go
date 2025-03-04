package controllers

import (
	"context"
	"fmt"
	"net/http"
	"server/database"
	platecode "server/models/plate-code"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbgorm"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RegisterCode struct {
	RegisterCode string `json:"registerCode"`
	RegisterCity string `json:"registerCity"`
	CodePosition *int   `json:"codePosition"`
	Note         string `json:"note"`
}

func CreateRegisterCode(ctx *gin.Context) {
	db := database.GetDB()
	registerCode := platecode.RegisterPlateCode{}
	regionCode := platecode.RegionPlateCode{}
	var newRegister RegisterCode

	region := ctx.Param("regionCode")

	// Ensure regionCode is provided
	if region == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "regionCode is required"})
		return
	}

	if err := ctx.ShouldBindJSON(&newRegister); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	newId := uuid.New()
	registerCode.IdRegisterCode = newId
	registerCode.RegisterCode = newRegister.RegisterCode
	registerCode.RegisterCity = newRegister.RegisterCity
	registerCode.Note = newRegister.Note
	registerCode.CodePosition = newRegister.CodePosition

	if err := crdbgorm.ExecuteTx(context.Background(), db, nil, func(tx *gorm.DB) error {
		if err := tx.Model(&regionCode).First(&regionCode, "region_code = ?", &region).Error; err != nil {
			return err
		}

		registerCode.IdRegionCode = regionCode.IdRegionCode

		result := tx.Where("register_code = ? AND register_city = ? AND id_region_code = ?", &newRegister.RegisterCode, &newRegister.RegisterCity, &regionCode.IdRegionCode).FirstOrCreate(&registerCode)

		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected == 0 {
			return fmt.Errorf("conflict: data already exists")
		}

		return nil
	}); err != nil {
		if err.Error() == "conflict: data already exists" {
			ctx.JSON(http.StatusConflict, gin.H{
				"error":   "Conflict when creating region",
				"message": "Data already exists",
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "New register successfully created",
	})
}

func GetRegisterCode(ctx *gin.Context) {
	db := database.GetDB()
	var RegisterCode []platecode.RegisterPlateCode

	// Query with ordering
	if err := db.Select("id_region_code, register_code, register_city, note, created_at").
		Order("id_region_code ASC").
		Order("register_code ASC").
		Where("id_status = ?", 1).
		Find(&RegisterCode).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	// Preallocate memory for better performance
	result := make([]gin.H, 0, len(RegisterCode))

	// Process data
	for _, rc := range RegisterCode {
		result = append(result, gin.H{
			"idRegister":    rc.IdRegisterCode,
			"registerCode":  rc.RegisterCode,
			"registerArea":  rc.RegisterCity,
			"registerNote":  rc.Note,
			"registerAdded": rc.CreatedAt,
		})
	}

	// Return JSON response
	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func GetRegisterCodeByRegionCode(ctx *gin.Context) {
	db := database.GetDB()
	var (
		RegisterCode []platecode.RegisterPlateCode
		result       = make([]gin.H, 0)
	)

	regionCode := ctx.Param("regionCode")

	// Ensure regionCode is provided
	if regionCode == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "regionCode is required"})
		return
	}

	// Corrected Query
	if err := db.Table("register_plate_codes AS register").
		Select("register.id_register_code, register.register_code, register.register_city, register.note, register.created_at").
		Joins("JOIN region_plate_codes AS region ON region.id_region_code = register.id_region_code").
		Where("register.id_status = ? AND region.region_code = ?", 1, regionCode).
		Order("register.register_code ASC").
		Find(&RegisterCode).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	// Process the result
	for _, rc := range RegisterCode {
		result = append(result, gin.H{
			"idRegister":    rc.IdRegisterCode,
			"registerCode":  rc.RegisterCode,
			"registerCity":  rc.RegisterCity,
			"registerNote":  rc.Note,
			"registerAdded": rc.CreatedAt,
		})
	}

	// Return the JSON response
	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
