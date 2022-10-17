PROMETHEUS_MANIFEST ?= prometheus.yml
GRAFANA_MANIFEST ?= grafana.yml

deploy:
	kubectl apply -f ${PROMETHEUS_MANIFEST}
	kubectl apply -f ${GRAFANA_MANIFEST}

clean:
	kubectl delete -f ${PROMETHEUS_MANIFEST}
	kubectl delete -f ${GRAFANA_MANIFEST}