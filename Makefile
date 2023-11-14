FRONT_END_BINARY=frontend-service
BROKER_BINARY=broker-service

up:
	@echo "Starting Docker images..."
	docker compose up -d
	@echo "Docker images started!"

up_build: build_broker
	@echo "Stopping docker images (if running...)"
	docker compose down
	@echo "Building (when required) and starting docker images..."
	docker compose up --build -d
	@echo "Docker images built and started!"

down:
	@echo "Stopping docker compose..."
	docker compose down
	@echo "Done!"

build_broker:
	@echo "Building broker binary..."
	cd broker-service && env GOOS=linux CGO_ENABLED=0 go build -o bin/${BROKER_BINARY} ./cmd/api
	@echo "Done!"

build_front:
	@echo "Building front end binary..."
	cd front-end && env CGO_ENABLED=0 go build -o bin/${FRONT_END_BINARY} ./cmd/web
	@echo "Done!"

start: build_front
	@echo "Starting front end"
	cd front-end && ./bin/${FRONT_END_BINARY} &

stop:
	@echo "Stopping front end..."
	@-pkill -SIGTERM -f "./bin/${FRONT_END_BINARY}"
	@echo "Stopped front end!"

.PHONY: up up_build down build_broker build_front start stop