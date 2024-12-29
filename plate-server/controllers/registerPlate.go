package controllers

import (
	"context"
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
	region := ctx.Param("region")
	registerCode := platecode.RegisterPlateCode{}
	regionCode := platecode.RegionPlateCode{}
	var newRegister RegisterCode

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

		if err := tx.Create(&registerCode).Error; err != nil {
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
		"message": "New register successfully created",
	})
}
