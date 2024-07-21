package handlers

import (
	"alertmanager/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Funciton to handle the initial route
func HandleAlerts(c *gin.Context) {
	//Webhook payload
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

// Alert processing with alerts we got from the alert manager
func processAlert(alert models.Alert) {
	//Check the alertname in the alert struct
	alertName, ok := alert.Labels["alertname"]
	if !ok {
		//Initiate the default alert
		handleDefaultAlert(alert)
		return
	}

	switch alertName {
	case "KubePodCrashLooping":
		//Initiate crashbackloopoff condition
		handlePodCrashLooping(alert)
	case "HighCPUUsage":
		//Initiate high cpu alert condition
		handleHighCPUUsage(alert)
	// Add more cases as needed for different alert types
	default:
		//Initiate the default alert
		handleDefaultAlert(alert)
	}
}
