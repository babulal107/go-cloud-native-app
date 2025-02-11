
build:
	docker build -t babulal107/go-cloud-native-app:latest .

list:
	docker images

running-container:
	docker ps

show-all-container:
	docker ps --all

create-network: # Create a network for the docker containers
	docker network create go-app

# --name=go-cloud-native-app : running container name
# -p 8081:8080 = specify port to access app on port 8081 and go-app running on port 8080
run:
	docker run --name=go-cloud-native-app --network=go-app -p 8080:8080 -it babulal107/go-cloud-native-app

# provides detailed information about Docker objects, such as containers, images, volumes, and networks, in JSON format.
inspect:
	docker inspect babulal107/go-cloud-app

# docker scout → Calls Docker Scout, a built-in security analysis tool in Docker.
# cves → Specifies that the scan should look for vulnerabilities (CVEs).
docker-scan:
	docker scout cves babulal107/go-cloud-app:latest

clear-stop-container:
	docker container prune -f

clear-unused-images:
	docker image prune -f
	## docker image prune -a

clean-up: ## This command removes: Stopped containers, Unused images, Unused networks, Build cache
	docker system prune

run-mongodb: # Run the mongodb container
	docker run --network=go-app --name mongodb -d -p 27017:27017 -v ~/opt/data:/data/db mongodb:latest

# using docker-compose file
up:
	docker compose up -d
	# docker compose up -d --build

down:
	docker compose down

check-logs:
	docker compose logs go_backend