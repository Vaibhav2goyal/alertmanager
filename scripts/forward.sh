kubectl port-forward svc/kube-prometheus-stack-prometheus 9090:9090 -n monitoring &
kubectl port-forward svc/kube-prometheus-stack-alertmanager 9093:9093 -n monitoring &
kubectl port-forward svc/alertmanager-webhook 8080:8080 -n monitoring &
kubectl port-forward svc/kube-prometheus-stack-grafana 3000:3000 -n monitoring
