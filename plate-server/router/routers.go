package router

import (
	"server/controllers"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	api := r.Group("/v1")

	auth := api.Group("/auth")
	{
		auth.POST("/register", controllers.UserRegister)
		auth.POST("/login", controllers.UserLogin)
	}

	vehicle := api.Group("/vehicle")
	{
		vehicle.GET("/category", controllers.GetVehicle)
		vehicle.Use(middleware.Auth())
		vehicle.POST("/engine", controllers.CreateVehicleEngine)
		vehicle.POST("/type", controllers.CreateVehicleType)
		vehicle.POST("/category", controllers.CreateVehicleCategory)
	}

	plateCode := api.Group("/plate-code")
	{
		plateCode.GET("/region", controllers.GetRegionCode)
		plateCode.GET("/register", controllers.GetRegisterCode)
		plateCode.GET("/register/:regionCode", controllers.GetRegisterCodeByRegionCode)
		plateCode.Use(middleware.Auth())
		plateCode.POST("/region", controllers.CreateRegionCode)
		plateCode.POST("/register/:regionCode", controllers.CreateRegisterCode)
	}

	checkData := api.Group("/check-data")
	{
		checkData.POST("/", controllers.CheckPlateData)
	}

	return r
}
