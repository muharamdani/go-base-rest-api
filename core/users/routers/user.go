package user

import (
	"github.com/gin-gonic/gin"
	"github.com/muharamdani/go-base-rest-api/core/users/controllers"
)

func Router(router *gin.Engine) {
	router.GET("/users", controllers.GetUser)
	// place other user routes here
}
