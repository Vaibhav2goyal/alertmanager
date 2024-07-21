package handlers

import (
	"alertmanager/actions"
	"alertmanager/enrichments"
	"alertmanager/models"
	"fmt"
)

func handleDefaultAlert(alert models.Alert) {
	fmt.Println("handleDefaultAlert")
	// Handle unknown or unrecognized alerts
	enrichments.RegisterEnrichment(enrichments.DefaultEnrichment{})

	enrichedData := enrichments.EnrichData(alert)

	actions.RegisterAction(actions.SlackAction{})

	actions.TakeAction(enrichedData)
}
