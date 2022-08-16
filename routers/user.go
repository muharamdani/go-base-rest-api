package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/muharamdani/go-base-rest-api/controllers"
)

func User(router *gin.Engine) {
	router.GET("/users", controllers.GetUser)
	// place other user routes here
}
