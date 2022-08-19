package controllers

import (
	"strconv"

	conn "github.com/muharamdani/go-base-rest-api/db"
	model "github.com/muharamdani/go-base-rest-api/pkg/users/models"
	repo "github.com/muharamdani/go-base-rest-api/pkg/users/repositories"
	request "github.com/muharamdani/go-base-rest-api/pkg/users/requests"
	"github.com/muharamdani/go-base-rest-api/utils"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	var user []model.Users
	var payload request.GetUser

	if err := c.Bind(&payload); err != nil {
		utils.ResponseBadRequest(c, "", err)
		return
	}

	if err := repo.FetchAll(conn.DB, &user, &payload); err != nil {
		utils.ResponseDatabaseError(c, err)
		return
	}

	utils.ResponseOK(c, "Get user data success", user)
}

func GetByID(c *gin.Context) {
	var user model.Users

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ResponseInternalServerError(c, err.Error())
		return
	}

	if err := repo.FetchByID(conn.DB, &user, id); err != nil {
		utils.ResponseDatabaseError(c, err)
		return
	}

	utils.ResponseOK(c, "Get user data success", user)
}

func Create(c *gin.Context) {
	var user model.Users

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ResponseUnprocessableEntity(c, "Invalid payload", err)
		return
	}

	if err := repo.Create(conn.DB, &user); err != nil {
		utils.ResponseDatabaseError(c, err)
		return
	}

	utils.ResponseOK(c, "Create user data success", user)
}

func Update(c *gin.Context) {
	var user model.Users

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ResponseInternalServerError(c, err.Error())
		return
	}

	if err := repo.FetchByID(conn.DB, &user, id); err != nil {
		utils.ResponseDatabaseError(c, err)
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ResponseUnprocessableEntity(c, "Invalid payload", err)
		return
	}

	if err := repo.Update(conn.DB, id, &user); err != nil {
		utils.ResponseDatabaseError(c, err)
		return
	}

	utils.ResponseOK(c, "Update user data success", user)
}

func Delete(c *gin.Context) {
	var user model.Users
	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ResponseInternalServerError(c, err.Error())
		return
	}

	if err := repo.FetchByID(conn.DB, &user, id); err != nil {
		utils.ResponseDatabaseError(c, err)
		return
	}

	if err := repo.Delete(conn.DB, id); err != nil {
		utils.ResponseDatabaseError(c, err)
		return
	}

	utils.ResponseOK(c, "Delete user data success", nil)
}