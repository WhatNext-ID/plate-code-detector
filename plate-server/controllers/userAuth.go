package controllers

import (
	"context"
	"net/http"
	"plate-server/database"
	"plate-server/helpers"
	"plate-server/models/user"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbgorm"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Register struct {
	UserName  string `json:"userName"`
	SecretKey string `json:"secretKey"`
	UserRole  string `json:"userRole"`
}

type Login struct {
	UserName  string `json:"userName"`
	SecretKey string `json:"secretKey"`
}

func UserRegister(ctx *gin.Context) {
	db := database.GetDB()
	userRegist := user.User{}
	userRole := user.UserRole{}
	var register Register

	if err := ctx.ShouldBindJSON(&register); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	newId := uuid.New()
	userRegist.IdUser = newId
	userRegist.UserName = register.UserName
	userRegist.SecretKey = register.SecretKey

	if err := crdbgorm.ExecuteTx(context.Background(), db, nil, func(tx *gorm.DB) error {
		if err := tx.Model(&userRole).First(&userRole, "role_name = ?", register.UserRole).Error; err != nil {
			return err
		}

		userRegist.IdRole = userRole.IdRole

		if err := tx.Create(&userRegist).Error; err != nil {
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
		"message": "Pengguna berhasil ditambahkan",
	})
}

func UserLogin(ctx *gin.Context) {
	db := database.GetDB()
	userData := user.User{}
	var userLogin Login

	if err := ctx.ShouldBindJSON(&userLogin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	err := db.Where("user_name = ?", userLogin.UserName).Take(&userData).Error

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Nama tidak ditemukan",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(userData.SecretKey), []byte(userLogin.SecretKey))

	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Kata sandi salah",
		})
		return
	}

	token := helpers.GenerateToken(userData.IdUser, userData.UserName)

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
