package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/muharamdani/gin-user-services/controllers"
)

func User(router *gin.Engine) {
	router.GET("/users", controllers.GetUser)
	// place other user routes here
}
