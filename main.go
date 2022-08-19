/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/muharamdani/go-base-rest-api/cmd"
	"github.com/muharamdani/go-base-rest-api/config"
	"github.com/muharamdani/go-base-rest-api/db"
	routers "github.com/muharamdani/go-base-rest-api/pkg"
)

func main() {
	cmd.Execute()

	// Initialize the database connection
	// and call env variable
	db.Connect("default")
	mode := config.Env("SET_MODE")
	port := config.Env("PORT")
	port = ":" + port

	// Set mode to debug|test|release based on the value of the environment variable
	gin.SetMode(mode)

	// Run gin instance
	app := gin.Default()
	app.SetTrustedProxies([]string{"127.0.0.1"})
	// Call for all router here
	routers.Router(app)

	// Listen and Server in 0.0.0.0:8080
	err := app.Run(port)
	if err != nil {
		return
	}
}
