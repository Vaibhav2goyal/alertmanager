package enrichments

import "alertmanager/models"

type Enrichment interface {
	Enrich(alert models.Alert) models.Alert
}

var enrichments []Enrichment

func EnrichData(alert models.Alert) models.Alert {
	for _, enrichment := range enrichments {
		alert = enrichment.Enrich(alert)
	}
	return alert
}

func RegisterEnrichment(e Enrichment) {
	enrichments = append(enrichments, e)
}
