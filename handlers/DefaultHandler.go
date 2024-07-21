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
	// Register defualt enrichment
	enrichments.RegisterEnrichment(enrichments.DefaultEnrichment{})
	//Enrich the data
	enrichedData := enrichments.EnrichData(alert)
	//Register slack action
	actions.RegisterAction(actions.SlackAction{})
	//Take the slack action
	actions.TakeAction(enrichedData)
}
