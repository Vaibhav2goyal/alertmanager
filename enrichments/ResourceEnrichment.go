package enrichments

import (
	"alertmanager/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type ResourceEnrichment struct {
	PrometheusURL string
}

// Prometheus response structure for getting the values
type PrometheusResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric map[string]string `json:"metric"`
			Value  []interface{}     `json:"value"` // Value is an array with two elements
		} `json:"result"`
	} `json:"data"`
}

func (e ResourceEnrichment) Enrich(alert models.Alert) models.Alert {
	podName := alert.Labels["pod"]
	namespace := alert.Labels["namespace"]
	if podName == "" || namespace == "" {
		return alert
	}
	//Getting the additonal mertics
	cpuUsage, memoryUsage := e.getPodMetrics(podName, namespace)
	alert.Labels["cpu_usage"] = cpuUsage
	alert.Labels["memory_usage"] = memoryUsage
	return alert
}

func (e ResourceEnrichment) getPodMetrics(podName, namespace string) (string, string) {
	//PromQL query to get the cpu usage
	cpuQuery := fmt.Sprintf(`sum(rate(container_cpu_usage_seconds_total{namespace="%s", pod="%s"}[5m]))`, namespace, podName)
	//PromQL query to get the memory usage
	memoryQuery := fmt.Sprintf(`sum(container_memory_usage_bytes{namespace="%s", pod="%s"})`, namespace, podName)
	//Getting CPU and memory usage with query created above
	cpuUsage := e.queryPrometheus(cpuQuery)
	memoryUsage := e.queryPrometheus(memoryQuery)

	return cpuUsage, memoryUsage
}

func (e ResourceEnrichment) queryPrometheus(query string) string {
	//Encoding the query properly
	encodedQuery := url.QueryEscape(query)
	fullURL := fmt.Sprintf("%s/api/v1/query?query=%s", e.PrometheusURL, encodedQuery)

	resp, err := http.Get(fullURL)
	if err != nil {
		fmt.Println("Error querying Prometheus:", err)
		return "N/A"
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Bad response from Prometheus: %s\n", resp.Status)
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("Response body:", string(body))
		return "N/A"
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading Prometheus response:", err)
		return "N/A"
	}

	// Decode JSON into the struct
	var promResponse PrometheusResponse
	if err := json.Unmarshal(body, &promResponse); err != nil {
		fmt.Println("Error unmarshalling JSON response:", err)
		return "N/A"
	}

	// Extract the value from the struct
	if len(promResponse.Data.Result) > 0 && len(promResponse.Data.Result[0].Value) > 1 {
		// Assuming the value is the second element in the array
		return fmt.Sprintf("%v", promResponse.Data.Result[0].Value[1])
	}

	return "N/A"
}
