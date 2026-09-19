package errorhandling

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServiceError struct {
	StatusCode int    `json:"-"`
	ErrorName  string `json:"error"`
	Message    string `json:"message"`
}

func (e *ServiceError) Error() string {
	return e.Message
}

func HandleServiceError(
	ctx *gin.Context,
	err error,
) {

	if serviceErr, ok := errors.AsType[*ServiceError](err); ok {
		ctx.JSON(serviceErr.StatusCode, serviceErr)
		return
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error":   "Internal Server Error",
		"message": err.Error(),
	})
}
