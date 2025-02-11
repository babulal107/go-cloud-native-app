## Go Cloud Native Application
 GoLang develop cloud native application with Restful API ( Docker+Kubernetes)

### Containerization Reference document: 
 Link: https://github.com/iam-veeramalla/MERN-docker-compose

### Create a network for the docker containers

`docker network create go_app_network`

### Build the Go app server

```sh
docker build -t babulal107/go-cloud-native-app:latest .
```

### Run the Go app server

`docker run --name=go_backend --network=go_app_network -d -p 8080:8080 -it babulal107/go-cloud-native-app`


### Verify Go app server running on localhost

Open your browser and type `http://localhost:8080`

## Using Docker Compose

Run
`docker compose up -d`
OR 
`docker compose up -d --build`

Stop:
`docker compose down`

Stop and remove volumes
`docker compose down -v`


### Checking Logs
`docker compose logs go_backend`

### Run the mongodb container

`docker run --network=demo --name mongodb -d -p 27017:27017 -v ~/opt/data:/data/db mongodb:latest`
