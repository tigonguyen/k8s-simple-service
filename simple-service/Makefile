APP_IMAGE_NAME ?= simple-service
TAG ?= v1.0.0
APP_MANIFEST ?= simple-service.yml

build:
	minikube image build -t ${APP_IMAGE_NAME}:${TAG} .

deploy:
	minikube kubectl -- apply -f ${APP_MANIFEST}

clean:
	minikube kubectl -- delete -f ${APP_MANIFEST}