package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/muharamdani/gin-user-services/routers/user"
)

func Main(router *gin.Engine) {
	// Call needed router here
	user.Router(router)
}
