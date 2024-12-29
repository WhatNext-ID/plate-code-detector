package controllers

import "github.com/gin-gonic/gin"

type DataCode struct {
	RegionCode        string  `json:"regionCode"`
	RegisterFirstCode *string `json:"firstCodeRegister"`
	RegisterLastCode  *string `json:"lastCodeRegister"`
}

func CheckCode(ctx *gin.Context) {

}
