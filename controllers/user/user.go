package user

import (
	"github.com/gin-gonic/gin"
	"github.com/muharamdani/gin-user-services/services"
	"github.com/muharamdani/gin-user-services/utils"
)

func GetUser(c *gin.Context) {
	data, err := services.GetUser(c)
	if err != nil {
		utils.ResponseInternalServerError(c, err.Error())
		return
	}
	message := "Get user data success"
	utils.ResponseOK(c, message, data)
}
