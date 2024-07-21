package handlers

import (
	"alertmanager/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleAlerts(c *gin.Context) {
	var payload models.WebhookPayload

	// Bind JSON payload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Process each alert in the payload
	for _, alert := range payload.Alerts {
		processAlert(alert)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Alerts processed"})
}

func processAlert(alert models.Alert) {
	alertName, ok := alert.Labels["alertname"]
	if !ok {
		handleDefaultAlert(alert)
		return
	}

	switch alertName {
	case "KubePodCrashLooping":
		handlePodCrashLooping(alert)
	case "HighCPUUsage":
		handleHighCPUUsage(alert)
	// Add more cases as needed for different alert types
	default:
		handleDefaultAlert(alert)
	}
}
