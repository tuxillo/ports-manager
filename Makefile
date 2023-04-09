clean:
	rm -fr vendor
	rm -f ports-manager

tidy:
	go mod tidy

vendor:
	go mod vendor

build: clean vendor tidy
	go build -o ports-manager cmd/manager.go
