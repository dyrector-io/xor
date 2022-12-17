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