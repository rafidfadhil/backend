IMAGE_NAME=rubenadi/bic-final-project-be:0.1.4

build:
	@go build -o bin/api cmd/api/main.go

run: build
	@./bin/api

docker.up:
	@docker-compose up -d

docker.down:
	@docker-compose down

docker.build:
	@docker build -t $(IMAGE_NAME) .

docker.push:
	@docker push $(IMAGE_NAME)

wire.auth:
	@wire ./internal/auth/di

wire.operational:
	@wire ./internal/operational/di