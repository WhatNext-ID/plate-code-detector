package router

import (
	"os"
	"plate-server/controllers"
	"plate-server/middleware"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.DebugMode // default to "debug"
	}

	gin.SetMode(ginMode)

	r := gin.Default()

	origins := os.Getenv("ORIGIN")
	raw := strings.Split(origins, ",")
	var allowedOrigins []string
	for _, o := range raw {
		allowedOrigins = append(allowedOrigins, strings.TrimSpace(o))
	}

	r.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
