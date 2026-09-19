package router

import (
	"os"
	"plate-server/middleware"
	"plate-server/modules/auth"
	checkplate "plate-server/modules/checkPlate"
	regionplate "plate-server/modules/regionPlate"
	registerplate "plate-server/modules/registerPlate"
	"plate-server/modules/vehicle"
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

	apiEndpoint := r.Group("/v1")

	authenticationEndpoint := apiEndpoint.Group("/auth")
	{
		authenticationEndpoint.POST("/register", auth.UserRegister)
		authenticationEndpoint.POST("/login", auth.UserLogin)
	}

	vehicleEndpoint := apiEndpoint.Group("/vehicle")
	{
		vehicleEndpoint.GET("/category", vehicle.GetVehicle)

		vehicleEndpoint.Use(middleware.Auth())
		vehicleEndpoint.POST("/engine", vehicle.CreateVehicleEngine)
		vehicleEndpoint.POST("/type", vehicle.CreateVehicleType)
		vehicleEndpoint.POST("/category", vehicle.CreateVehicleCategory)
	}

	plateCodeEndpoint := apiEndpoint.Group("/plate-code")
	{
		plateCodeEndpoint.GET("/region", regionplate.GetRegionCode)
		plateCodeEndpoint.GET("/register", registerplate.GetRegisterCode)
		plateCodeEndpoint.GET("/register/:regionCode", registerplate.GetRegisterCodeByRegionCode)

		plateCodeEndpoint.Use(middleware.Auth())
		plateCodeEndpoint.POST("/region", regionplate.CreateRegionCode)
		plateCodeEndpoint.POST("/register/:regionCode", registerplate.CreateRegisterCode)
	}

	checkDataEndpoint := apiEndpoint.Group("/check-data")
	{
		checkDataEndpoint.POST("/", checkplate.CheckPlateData)
	}

	return r
}
