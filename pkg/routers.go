package pkg

import (
	"github.com/gin-gonic/gin"
	users "github.com/muharamdani/go-base-rest-api/pkg/users/routers"
)

func Main(router *gin.Engine) {
	// Call needed router here
	users.Router(router)
}
