package router

import (
	"server/controller"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	statusKendaraan := r.Group("/status-kendaraan")
	{
		statusKendaraan.POST("/", controller.PostStatusKendaraan)
		statusKendaraan.PATCH("/:id", controller.UpdateStatusKendaraan)
		statusKendaraan.GET("/", controller.GetAllStatusKendaraan)
	}

	return r
}
