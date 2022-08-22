package tests

import (
	"net/http"
	"testing"

	"github.com/muharamdani/go-base-rest-api/pkg/users/routers"
	"github.com/muharamdani/go-base-rest-api/utils"

	"github.com/gin-gonic/gin"
)


func TestGetUsers(t *testing.T) {
	engine := gin.Default()
	routers.Router(engine)
	e := utils.PrepareTest(t, engine)
	
	e.GET("/users").
		Expect().
		Status(http.StatusOK).JSON().Object().ValueEqual("message", "Get user data success")

	defer utils.CleanUp("users")
}

func TestInsertUser(t *testing.T) {
	engine := gin.Default()
	routers.Router(engine)
	e := utils.PrepareTest(t, engine)
	
	resp := e.POST("/users").
		WithJSON(map[string]interface{}{
			"first_name": "Muhamad",
			"last_name": "Ramdani",
			"username": "muharamdani",
			"phone_number": "812317641",
			"email": "muharamdani@gmail.com",
			"address": "Rumah no 91",
		}).Expect().Status(http.StatusOK).JSON().Object()

	resp.Value("data").Object().ValueEqual("first_name", "Muhamad")

	defer utils.CleanUp("users")
}