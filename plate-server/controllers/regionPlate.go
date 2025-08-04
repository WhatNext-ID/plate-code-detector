package controllers

import (
	"net/http"
	"plate-server/database"
	platecode "plate-server/models/plate-code"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RegionCode struct {
	Code         string  `json:"code"`
	Area         string  `json:"area"`
	Note         *string `json:"note"`
	CodePosition string  `json:"position"`
}

func CreateRegionCode(ctx *gin.Context) {
	db := database.GetDB()
	regionCode := platecode.RegionPlateCode{}
	var newRegion RegionCode

	if err := ctx.ShouldBindJSON(&newRegion); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	newId := uuid.New()
	regionCode.IdRegionCode = newId
	regionCode.RegionCode = newRegion.Code
	regionCode.RegionArea = newRegion.Area
	regionCode.Note = newRegion.Note

	result := db.Where("region_code = ? AND region_area = ?", &newRegion.Code, &newRegion.Area).FirstOrCreate(&regionCode)

	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusConflict, gin.H{
			"error":   "Conflict when create region",
			"message": "Data already exist",
		})
		return
	}

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": result.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "New region successfully created",
	})
}

func GetRegionCode(ctx *gin.Context) {
	db := database.GetDB()
	var RegionCode []platecode.RegionPlateCode

	// Query with ordering
	if err := db.Select("id_region_code, region_code, region_area, note, created_at").
		Order("created_at DESC").
		Where("id_status = ?", 1).
		Find(&RegionCode).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	// Preallocate memory for efficiency
	result := make([]gin.H, 0, len(RegionCode))

	// Transform data
	for _, rc := range RegionCode {
		result = append(result, gin.H{
			"idRegion":    rc.IdRegionCode,
			"regionCode":  rc.RegionCode,
			"regionArea":  rc.RegionArea,
			"regionNote":  rc.Note,
			"regionAdded": rc.CreatedAt,
		})
	}

	// Return JSON response
	ctx.JSON(http.StatusOK, gin.H{"data": result})
}
