package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func ResponseOK(c *gin.Context, msg string, data interface{}) {
	result := Response{
		Status:  true,
		Message: msg,
		Data:    data,
	}
	c.JSON(http.StatusOK, result)
}

func ResponseInternalServerError(c *gin.Context, msg string) {
	result := Response{
		Status:  false,
		Message: msg,
		Data:    nil,
	}
	c.JSON(http.StatusInternalServerError, result)
}

func ResponseBadRequest(c *gin.Context, msg string, err error) {
	result := Response{
		Status:  false,
		Message: msg,
		Data:    err,
	}
	c.JSON(http.StatusBadRequest, result)
}

func ResponseUnauthorized(c *gin.Context, msg string) {
	result := Response{
		Status:  false,
		Message: msg,
		Data:    nil,
	}
	c.JSON(http.StatusUnauthorized, result)
}

func ResponseForbidden(c *gin.Context, msg string) {
	result := Response{
		Status:  false,
		Message: msg,
		Data:    nil,
	}
	c.JSON(http.StatusForbidden, result)
}

func ResponseUnprocessableEntity(c *gin.Context, msg string, err error) {
	validationResult := GetValidationResult(err)
	result := Response{
		Status:  false,
		Message: msg,
		Data:    validationResult,
	}
	c.JSON(http.StatusUnprocessableEntity, result)
}
