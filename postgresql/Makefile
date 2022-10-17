DB_IMAGE_NAME ?= postgres-ssl
TAG ?= v1.0.0
DB_MANIFEST ?= postgresql.yml

build:
	docker build -t ${DB_IMAGE_NAME}:${TAG} .

deploy:
	kubectl apply -f ${DB_MANIFEST}

clean:
	kubectl delete -f ${DB_MANIFEST}