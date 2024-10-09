package controllers

import (
	"net/http"
	"server/database"
	"server/helpers"
	"server/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AdminRegister(ctx *gin.Context) {
	db := database.GetDB()
	UserAdmin := models.MAdmin{}

	if err := ctx.ShouldBindJSON(&UserAdmin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	newId := uuid.New()
	UserAdmin.IdAdmin = newId

	if err := db.Create(&UserAdmin).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User Berhasil Dibuat",
		"nama":    UserAdmin.NamaAdmin,
		"id":      UserAdmin.IdAdmin,
	})
}

func AdminLogin(ctx *gin.Context) {
	db := database.GetDB()
	User := models.MAdmin{}
	password := ""

	if err := ctx.ShouldBindJSON(&User); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	password = User.SandiAdmin

	err := db.Where("nama_admin = ?", User.NamaAdmin).Take(&User).Error

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid name or password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.SandiAdmin), []byte(password))

	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email or password",
		})
		return
	}

	token := helpers.GenerateToken(User.IdAdmin, User.NamaAdmin)

	if token == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Failed to generate token",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token":   token,
		"message": "Selamat anda berhasil masuk",
	})
}
