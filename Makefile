.PHONY: cli
cli: 
	cd api/cmd/cli && \
	go run .; \
	cd -


.PHONY: server
server: 
	cd api/cmd/server && \
	go run .; \
	cd -


.PHONY: compile
compile:
	cd api/cmd/server && \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./api.bin

.PHONY: build-api
build-api:
	cd api && \
	docker build -t github.com/dyrector-io/xor/api:v1 .
	