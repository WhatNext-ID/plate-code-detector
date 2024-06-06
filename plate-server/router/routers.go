package router

import (
	"server/controllers"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userAdmin := r.Group("/user")
	{
		userAdmin.POST("/register", controllers.AdminRegister)
		userAdmin.POST("/login", controllers.AdminLogin)
	}

	statusKendaraan := r.Group("/status-kendaraan")
	{
		statusKendaraan.GET("/", controllers.GetAllStatusKendaraan)
		statusKendaraan.Use(middleware.Auth())
		statusKendaraan.POST("/", controllers.PostStatusKendaraan)
		statusKendaraan.PATCH("/:id", controllers.UpdateStatusKendaraan)
	}

	kodeWilayah := r.Group("/kode-wilayah")
	{
		kodeWilayah.POST("/", controllers.PostKodeWilayah)
	}

	return r
}
