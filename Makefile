infra:
	minikube start
	minikube kubectl
	
build:
	cd simple-service && $(MAKE) build
	cd postgresql && $(MAKE) build

deploy:
	cd simple-service && $(MAKE) deploy
	cd postgresql && $(MAKE) deploy
	cd monitoring && $(MAKE) deploy

clean: 
	cd simple-service && $(MAKE) clean
	cd postgresql && $(MAKE) clean
	cd monitoring && $(MAKE) clean