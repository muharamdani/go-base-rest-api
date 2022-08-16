package services

import (
	"strconv"

	conn "github.com/muharamdani/go-base-rest-api/core/db"
	"github.com/muharamdani/go-base-rest-api/core/models"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) (interface{}, error) {
	var users []models.Users
	perPageStr := c.DefaultQuery("per_page", "10")

	// convert perPage to int
	perPageInt, err := strconv.Atoi(perPageStr)
	if err != nil {
		return nil, err
	}
	conn.DB.Limit(perPageInt).Find(&users)
	return users, nil
}
