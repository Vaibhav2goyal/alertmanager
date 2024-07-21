package main

import (
	"alertmanager/config"
	"alertmanager/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	r := gin.Default()
	r.POST("/alerting", handlers.HandleAlerts)
	r.Run(":8080")
}
