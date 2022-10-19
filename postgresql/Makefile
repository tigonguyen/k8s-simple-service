DB_IMAGE_NAME ?= postgres-ssl
TAG ?= v1.0.0
DB_MANIFEST ?= postgresql.yml

build:
	minikube image build -t ${DB_IMAGE_NAME}:${TAG} .

deploy:
	minikube kubectl -- apply -f ${DB_MANIFEST}

clean:
	minikube kubectl -- delete -f ${DB_MANIFEST}