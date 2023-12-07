full : test build run

test:
	@ echo "Running test"
	@ go test -race -v ./...
	@ echo "Test passed"

build: ## Builds binary
	@ echo "Vendoring"
	@ go mod vendor
	@ echo "Building aplication... "
	@ go build \
		-o ./app/ \
		./app/library.go
	@ echo "done"

build-race: ## Builds binary (with -race flag)
	@ echo "Vendoring"
	@ go mod vendor
	@ echo "Building aplication with race flag... "
	@ go build \
		-race      \
		-o ./app/ \
		./app/library.go
	@ echo "done"

run :
	@ ./app/library