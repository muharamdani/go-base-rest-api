package routers

import (
	"github.com/gin-gonic/gin"
)

func Main(router *gin.Engine) {
	// Call needed router here
	User(router)
}
