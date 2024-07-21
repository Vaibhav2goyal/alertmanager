package handlers

import (
	"alertmanager/actions"
	"alertmanager/enrichments"
	"alertmanager/models"
	"fmt"
	"os"
)

func handlePodCrashLooping(alert models.Alert) {

	fmt.Println("handlePodCrashLooping")
	// Register enrichment strategies specific to Pod Crash Looping alerts
	enrichments.RegisterEnrichment(enrichments.ResourceEnrichment{PrometheusURL: os.Getenv("PROMETHEUS_URL")})

	enrichments.RegisterEnrichment(enrichments.DefaultEnrichment{})

	// Enrich alert data
	enrichedData := enrichments.EnrichData(alert)

	// Register actions specific to Pod Crash Looping alerts
	actions.RegisterAction(actions.SlackAction{})

	// Take all the  registered actions
	actions.TakeAction(enrichedData)
}

func handleHighCPUUsage(alert models.Alert) {
	fmt.Println("handleHighCPUUsage")
	// Register enrichment strategies specific to High CPU Usage alerts
	enrichments.RegisterEnrichment(enrichments.DefaultEnrichment{})

	enrichments.RegisterEnrichment(enrichments.ResourceEnrichment{PrometheusURL: os.Getenv("PROMETHEUS_URL")})
	// Enrich alert data
	enrichedData := enrichments.EnrichData(alert)

	// Register actions specific to High cpu alerts
	actions.RegisterAction(actions.SlackAction{})

	// Take all the  registered actions
	actions.TakeAction(enrichedData)
}
