PROMETHEUS_MANIFEST ?= prometheus.yml
GRAFANA_MANIFEST ?= grafana.yml

deploy:
	minikube kubectl -- apply -f ${PROMETHEUS_MANIFEST}
	minikube kubectl -- apply -f ${GRAFANA_MANIFEST}

clean:
	minikube kubectl -- delete -f ${PROMETHEUS_MANIFEST}
	minikube kubectl -- delete -f ${GRAFANA_MANIFEST}