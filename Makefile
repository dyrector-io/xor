.PHONY: cli
cli:
	cd api/cmd/cli && go run .

.PHONY: server
server:
	cd api/cmd/server && \
	go run .

.PHONY: compile
compile:
	cd api/cmd/server && \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./api.bin

.PHONY: build-api
build-api: compile
	docker-compose build api

.PHONY: build-ui
build-ui:
	docker-compose build ui

.PHONY: all
all: build-api build-ui
	docker-compose push
