package auth

import (
	"net/http"
	errorhandling "plate-server/utils/error-handling"

	"github.com/gin-gonic/gin"
)

func UserRegister(ctx *gin.Context) {
	var registerUser Register

	if err := ctx.ShouldBindJSON(&registerUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	result, err := register(registerUser)
	if err != nil {
		errorhandling.HandleServiceError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": result,
	})
}

func UserLogin(ctx *gin.Context) {
	var userLogin Login

	if err := ctx.ShouldBindJSON(&userLogin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	result, err := login(userLogin)
	if err != nil {
		errorhandling.HandleServiceError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}
