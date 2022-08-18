package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	conn "github.com/muharamdani/go-base-rest-api/core/db"
	model "github.com/muharamdani/go-base-rest-api/core/models/user"
	repo "github.com/muharamdani/go-base-rest-api/core/repositories/user"
	request "github.com/muharamdani/go-base-rest-api/core/requests/user"
	"github.com/muharamdani/go-base-rest-api/utils"
)


func GetUser(c *gin.Context) {
	var result []model.Users

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		utils.ResponseInternalServerError(c, err.Error())
		return
	}

	payload := request.FetchAllUser{
		Limit: limit,
	}

	if err := repo.FetchAll(conn.DB, &result, &payload); err != nil {
		utils.ResponseInternalServerError(c, err.Error())
		return
	}

	message := "Get user data success"
	utils.ResponseOK(c, message, result)
}