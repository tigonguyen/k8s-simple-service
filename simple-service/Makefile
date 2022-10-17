APP_IMAGE_NAME ?= simple-service
TAG ?= v1.0.0
APP_MANIFEST ?= simple-service.yml

build:
	docker build -t ${APP_IMAGE_NAME}:${TAG} .

deploy:
	kubectl apply -f ${APP_MANIFEST}

clean:
	kubectl delete -f ${APP_MANIFEST}