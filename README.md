# Alert Manager

## Overview

This project is an Alert Manager that handles incoming alerts with any payload in an JSON format, enriches them with additional data, and takes actions based on the data. New types of handlers, enrichment and actions can be easily added with the interface.

## Structure

alertmanager/
├── main.go
├── handlers/
│   ├── Handler.go
│   ├── DeafaultHandler.go
│   ├── PodHandler.go
├── enrichments/
│   ├── DefaultEnrichments.go
│   ├── Enrichments.go
│   ├── ResourceEnrichments.go
├── models/
│   └── AlertModel.go
├── config/
│   └── config.go
├── scripts/
│   ├── kube-prometheus-stack
│   ├── deployment.yaml
│   ├── start.sh
├── go.mod
└── go.sum
└── Dockerfile


## Setup

1. Clone the repository:
    ```sh
    git clone <repository_url>
    cd alertmanager
    ```

2. Set environment variables for Slack integration and Prometheus url in .env file :
    
    SLACK_API_TOKEN=your-slack-token
    SLACK_CHANNEL_ID=your-slack-channel-id
    PrometheusURL=http://kube-prometheus-stack-prometheus:9090

3. Install kind cluster(Macos) and create a cluster:
    ```sh
    brew install kind

    kind create cluster --name alertmanager  

    kubectl config use-context kind-alertmanager
    ```

4. Build a Docker Image for alertmanager-webhook:
    ```sh

    docker build -t alertmanager-webhook . 
    
    ```

5. Load the docker image to the kind cluster:
    ```sh
    kind load docker-image alertmanager-webhook:latest --name alertmanager
    ```

6. Apply the alertmanager-webhook manifest to deploy the application:
    ```sh
    kubectl apply -f ./scripts/deployment.yaml  
    ```
7. Install kube-prometheus-stack with helm charts:
    ```sh
    helm upgrade --install kube-prometheus-stack  -f ./scripts/kube-prometheus-stack/values.yaml ./scripts/kube-prometheus-stack/ 
    ```

    


## Extending the System


### Adding a New Handler

1. Implement the new `Handler` case in Handler.go:
    ```go
	switch alertName {
	case "KubePodCrashLooping":
		handlePodCrashLooping(alert)
	case "HighCPUUsage":
		handleHighCPUUsage(alert)

	// Add more cases as needed for different alertname types
    case "MyCase":
		handleMyCase(alert)



    default:
		handleDefaultAlert(alert)
	}

    
    ```

2. Create a new handler file according to the need:
    ```go
    import (
	"alertmanager/actions"
	"alertmanager/enrichments"
	"alertmanager/models"
    )
    
    func handleMyCase(alert models.Alert) {

	fmt.Println("handleMyCase")
	// Register enrichment strategies specific to MyCase alerts
    enrichments.RegisterEnrichment(enrichments.MyNewEnrichment{})

	// Enrich alert data
	enrichedData := enrichments.EnrichData(alert)

	// Register actions specific to MyCase alerts
	actions.RegisterAction(actions.MyCaseAction{})

	// Take all the  registered actions
	actions.TakeAction(enrichedData)
    }      
   
    ```

### Adding a New Enrichment

1. Implement the `Enrichment` interface:
    ```go
    package enrichments

    //Import models
    import "alertmanager/models"
    type MyNewEnrichment struct{}

    func (e MyNewEnrichment) Enrich(alert models.Alert) models.Alert {
        // Enrichment logic
        return alert
    }
    ```

2. Register the new enrichment in the every handler you want to use:
    ```go
    import "alertmanager/enrichments"

         // other code...
        enrichments.RegisterEnrichment(enrichments.MyEnrichment{})
        // other code...
   
    ```

### Adding a New Action

1. Implement the `Action` interface:
    ```go
    package actions
    //Import models
    import "alertmanager/models"

    type MyNewAction struct{}

    func (a MyNewAction) Execute(alert Alert) {
        // Action logic
    }
    ```

2. Register the new action in the handler file:
    ```go
    import "alertmanager/actions"

        // other code...
        actions.RegisterAction(actions.MyAction{})
        // other code...
    
    ```

## Testing

You can test the implementation by sending a POST request to the `/alerting` endpoint with a sample alert payload.


