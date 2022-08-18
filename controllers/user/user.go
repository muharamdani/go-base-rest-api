package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	conn "github.com/muharamdani/go-base-rest-api/core/db"
	"github.com/muharamdani/go-base-rest-api/core/models"
	"github.com/muharamdani/go-base-rest-api/core/repositories/user"
	"github.com/muharamdani/go-base-rest-api/utils"
)

type getUserPayload struct {
	Limit int
}

func GetUser(c *gin.Context) {
	var result []models.Users

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		utils.ResponseInternalServerError(c, err.Error())
		return
	}

	payload := getUserPayload{
		Limit: limit,
	}

	if err := user.FetchAll(conn.DB, &result, payload); err != nil {
		utils.ResponseInternalServerError(c, err.Error())
		return
	}

	message := "Get user data success"
	utils.ResponseOK(c, message, result)
}
