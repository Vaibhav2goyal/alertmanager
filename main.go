package main

import (
	"alertmanager/config"
	"alertmanager/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	//Load the env
	config.LoadEnv()

	//Starting the gin server
	r := gin.Default()
	//Handling the /alerting route
	r.POST("/alerting", handlers.HandleAlerts)
	r.Run(":8080")
}
