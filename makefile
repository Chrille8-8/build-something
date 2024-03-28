# Some code of this makefile is created with help of chatGPT
# Author: Christoffer Persson 


# Names/tags for the microservices
SERVER:= server
ENCRYPTION := encryption
DATABASE := database
# Name of service
SERVICENAME := server-service

#
# ----------------- Start minikube and the micro-services -----------------
#
start_service:
	minikube start
	kubectl apply -f micro_services.yml
	minikube service $(SERVICENAME)

#
# ----------------- Build and push to dockerhub ------------------- 
#
build_server:
# Build the service that the user interacts with
# The if statement below is to handle when the service does not exist, just to make to rule work and not stop.
	@cd server && \
	if docker ps -a --format '{{.Names}}' | grep -Eq "^${SERVER}\$$"; then \
    	docker rm ${SERVER}; \
		docker rmi ${SERVER}:latest \
	else \
    	echo "Container ${SERVER} does not exist."; \
	fi
#	@docker build -t $(SERVER) -f server/Dockerfile server
	@docker build --no-cache -t $(SERVER) -f server/Dockerfile server
	@docker tag $(SERVER):latest ch12388/build_something:$(SERVER)

# Push the image to dockerhub
push_server:	
	@docker push ch12388/build_something:$(SERVER)


build_encryption:
# Build and push the encryption service
	@cd encryption && \
	if docker ps -a --format '{{.Names}}' | grep -Eq "^${ENCRYPTION}\$$"; then \
    	docker rm ${ENCRYPTION}; \
		docker rmi ${ENCRYPTION}:latest \
	else \
    	echo "Container ${ENCRYPTION} does not exist."; \
	fi
	@docker build -t $(ENCRYPTION) -f encryption/Dockerfile encryption
	@docker tag $(ENCRYPTION):latest ch12388/build_something:$(ENCRYPTION)

push_encryption:
	@docker push ch12388/build_something:$(ENCRYPTION)

build_database:
	@cd database && \
	if docker ps -a --format '{{.Names}}' | grep -Eq "^${DATABASE}\$$"; then \
    	docker rm ${DATABASE}; \
		docker rmi ${DATABASE}:latest \
	else \
    	echo "Container ${DATABASE} does not exist."; \
	fi
	@docker build -t $(DATABASE) -f database/Dockerfile database
	@docker tag $(DATABASE):latest ch12388/build_something:$(DATABASE)

push_database:
	@docker push ch12388/build_something:$(DATABASE)

# Build and push all three microservices
build_and_push_all: build_server build_encryption build_database push_server push_encryption push_database 
build: build_server build_encryption build_database 



#------------------------- GOLANG STUFF ---------------------
#
# Use this to build and run all services using golang
go_build_server:
	@cd server && \
	go build

go_run_server:
	@cd server && ./server

go_run_encryption:
	@cd encryption && ./encrypt

go_run_database:
	@cd database && ./db

go_build_encryption:
	@cd encryption && \
	go build

go_build_database:
	@cd database && \
	go build

go_build_all: go_build_server go_build_encryption go_build_database

#-------------------------- DOCKER & KUBERNETES STUFF ---------------------
#
#
create_docker_network:
	docker network create microservices-network

docker_run_server_on_network:
	docker build -t server-image server
	docker run -d --name server-container --network microservices-network -p 5000:5000 server-image

docker_run_encryption_on_network:
	docker build -t encrypt-image encryption
	docker run -d --name encrypt-container --network microservices-network -p 5001:5001 encrypt-image

docker_run_database_on_network:
	docker build -t database-image database
	docker run -d --name database-container --network microservices-network -p 5002:5002 database-image
#
# ----------------- Lsit docker & kubernetse stuff -----------------
#
docker_list:
	@docker images
	@docker ps -a 

kube_list:
	kubectl get all
	kubectl get pods
	kubectl get services
	kubectl get deployments
	kubectl get nodes

#
# ----------------- remove docker & kubernetes stuff -----------------
#
remove_all_containers:
	docker stop $(shell docker ps -aq)
	docker rm $(shell docker ps -aq)


remove_all_images:
	docker image prune -a


remove_kubernetes_resources:
	kubectl delete all --all --all-namespaces
#	kubectl delete deployment microservices --ignore-not-found=true
#	kubectl delete service microservices --ignore-not-found=true

remove_all: remove_kubernetes_resources remove_all_containers remove_all_images

#
# ----------------- Start the microservices (as binaries)----------------- 
#
start_server_no_kube: kill_process_on_port_5000
	@cd server && go build && ./server
	
start_encrypt_no_kube: kill_process_on_port_5001
	@cd encryption && go build && ./encrypt

start_database_no_kube: kill_process_on_port_5002
	@cd database && go build && ./db


#
# ----------------- Free ports ----------------- 
# This rule will kill the processes that are using the ports 5000, 5001 and 5002
kill_process_on_port_5000:
# Kill Server
	@PID_0=$$(sudo lsof -ti :5000); \
	if [ -n "$$PID_0" ]; then \
		echo "Process running on port 5000 with PID: $$PID_0. Killing it..."; \
		sudo kill $$PID_0; \
	else \
		echo "No process running on port 5000."; \
	fi

kill_process_on_port_5001:
# Kill encryption
	@PID_1=$$(sudo lsof -ti :5001); \
	if [ -n "$$PID_1" ]; then \
		echo "Process running on port 5001 with PID: $$PID_1. Killing it..."; \
		sudo kill $$PID_1; \
	else \
		echo "No process running on port 5001."; \
	fi

kill_process_on_port_5002:
# Kill database
		@PID_2=$$(sudo lsof -ti :5002); \
	if [ -n "$$PID_2" ]; then \
		echo "Process running on port 5002 with PID: $$PID_2. Killing it..."; \
		sudo kill $$PID_2; \
	else \
		echo "No process running on port 5002."; \
	fi

