package enrichments

import "alertmanager/models"

type DefaultEnrichment struct {
}

// Default enrichment
func (e DefaultEnrichment) Enrich(alert models.Alert) models.Alert {
	alert.Labels["enriched"] = "true"
	return alert
}
