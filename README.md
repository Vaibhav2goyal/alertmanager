# Alert Manager

## Overview

This project is an Alert Manager that handles incoming alerts with any payload in an JSON format, enriches them with additional data, and takes actions based on the data. New types of handlers, enrichment and actions can be easily added with the interface.

## 📜 Prerequisites
To easily use it you should have docker, kubectl and helm installed on your machine to run the application.

---

## 📄 Code Structure
```
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

```

- **main.go**

This contains main function which initialize the gin server. We load the env here and handle the routes.

- **handlers/Handler.go**

This file is used to intiate the handler package which handles the route /alerting and processes the webhook payload then process the alert according to its alertName. 

- **enrichments/Enrichment.go**

This file is used to initiates the enrichment package and uses enrichment interface to process other enrichments with the Enrich function

- **actions/action.go**

This file is being used to the actions package which after enriching data in the previous step can send the alert to multiple notification channels. We need to handle which action we want to use in handler file.

- **handlers package**

This package is used to handle additional alerts that need to be configured. Check Adding a New Handler section to add a new handler alert

- **enrichments package**

This package is used to add enrichments additional data to the alert. Check Adding a New Enrichment section to add a new Enrichment.

- **actions package**

This package is used to add actions for the alerts. Check Adding a New Alert section to add a new Alert

---

## Setup

### 1. Clone the repository:
    ```sh
    git clone git@github.com:Vaibhav2goyal/alertmanager.git
    cd alertmanager
    ```

### 2. Create a .env file
 Set environment variables for Slack integration and Prometheus url in .env file :

    ```
    SLACK_API_TOKEN=your-slack-token
    SLACK_CHANNEL_ID=your-slack-channel-id
    PrometheusURL=http://kube-prometheus-stack-prometheus:9090

    ```

### 3. Install kind cluster(MacOs) and create a cluster:
You can also use any other cluster of your choice but kind is recommended for now.        

    ```sh
    brew install kind

    kind create cluster --name alertmanager  

    kubectl config use-context kind-alertmanager
    ```

### 4. Build a Docker Image for alertmanager-webhook:
    ```sh

    docker build -t alertmanager-webhook . 
    
    ```

### 5. Load the Image
Load the docker image to the kind cluster:
   
    ```sh
    kind load docker-image alertmanager-webhook:latest --name alertmanager
    ```

### 6. Apply the alertmanager-webhook manifest to deploy the application:
    ```sh
    kubectl apply -f ./scripts/deployment.yaml  -n monitoring
    ```
### 7. Install kube-prometheus-stack with helm charts:
Optional - This is for observability of the the service and also fires the alerts for the service
    ```sh
    helm upgrade --install kube-prometheus-stack  -f ./scripts/kube-prometheus-stack/values.yaml ./scripts/kube-prometheus-stack/ -n monitoring
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


