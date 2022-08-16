package user

import (
	"github.com/gin-gonic/gin"
	"github.com/muharamdani/gin-user-services/controllers/user"
)

func Router(router *gin.Engine) {
	router.GET("/users", user.GetUser)
	// place other user routes here
}
