package main

import (
	"github.com/muharamdani/go-base-rest-api/config"
	routers "github.com/muharamdani/go-base-rest-api/core"
	"github.com/muharamdani/go-base-rest-api/core/db"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	// and call env variable
	db.Connect()
	mode := config.Env("SET_MODE")
	port := config.Env("PORT")
	port = ":" + port

	// Set mode to debug|test|release based on the value of the environment variable
	gin.SetMode(mode)

	// Run gin instance
	app := gin.Default()
	app.SetTrustedProxies([]string{"127.0.0.1"})
	// Call for all router here
	routers.Main(app)

	// Listen and Server in 0.0.0.0:8080
	err := app.Run(port)
	if err != nil {
		return
	}
}
