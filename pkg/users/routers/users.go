package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/muharamdani/go-base-rest-api/pkg/users/controllers"
)

func Router(router *gin.Engine) {
	router.GET("/users", controllers.Get)
	router.GET("/users/:id", controllers.GetByID)
	router.POST("/users", controllers.Create)
	router.PUT("/users/:id", controllers.Update)
	router.DELETE("/users/:id", controllers.Delete)
}
