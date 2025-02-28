package controllers

import "github.com/gin-gonic/gin"

type DataCode struct {
	RegionCode        string  `json:"regionCode"`
	RegisterFirstCode *string `json:"firstCodeRegister"`
	RegisterLastCode  *string `json:"lastCodeRegister"`
}

// Check Vehicle Type or Status by Plate Color
func CheckVehicleStatus(ctx *gin.Context) {

}

// Check Vehicle Region
func CheckVehicleRegion(ctx *gin.Context) {

}

// Check Vehicle Register Area
func CheckVehicleRegisterArea(ctx *gin.Context) {

}

// Check Plate Data
func CheckPlateData(ctx *gin.Context) {
	// Check Vehicle Status
	// Check Region Area
	// Get Code Position
	// Check Register Area
	// Return Possibility Match Data About The Plate
}
