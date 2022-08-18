package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/muharamdani/go-base-rest-api/core/users/routers"
)

func Main(router *gin.Engine) {
	// Call needed router here
	user.Router(router)
}
