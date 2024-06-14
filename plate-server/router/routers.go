package router

import (
	"server/controllers"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	api := r.Group("/v1")

	userAdmin := api.Group("/user")
	{
		userAdmin.POST("/register", controllers.AdminRegister)
		userAdmin.POST("/login", controllers.AdminLogin)
	}

	statusKendaraan := api.Group("/status-kendaraan")
	{
		statusKendaraan.GET("/", controllers.GetAllStatusKendaraan)
		statusKendaraan.Use(middleware.Auth())
		statusKendaraan.POST("/", controllers.PostStatusKendaraan)
		statusKendaraan.PATCH("/:id", controllers.UpdateStatusKendaraan)
	}

	kodeWilayah := api.Group("/kode-wilayah")
	{
		kodeWilayah.GET("/:id", controllers.GetKodeWilayahById)
		kodeWilayah.GET("/", controllers.GetAllKodeWilayah)
		kodeWilayah.Use(middleware.Auth())
		kodeWilayah.POST("/", controllers.PostKodeWilayah)
		kodeWilayah.PATCH("/:id", controllers.UpdateKodeWilayah)
	}

	kodeRegistrasi := api.Group("/kode-registrasi")
	{
		kodeRegistrasi.POST("/", controllers.PostKodeRegister)
	}

	kodeRegistrasiKhusus := api.Group("/kode-khusus")
	{
		kodeRegistrasiKhusus.POST("/", controllers.PostSpecialRegisterCode)
	}

	checkCode := api.Group("/cek-kode")
	{
		checkCode.POST("/", controllers.CheckCode)
	}

	return r
}
