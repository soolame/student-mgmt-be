# Variables
BINARY_NAME=student-app
DOCKER_USER=soolame

.PHONY: hello build run env generate-schema migrate-up migrate-down create-admin lint test clean docker-build

hello:
	echo "Hello"

# NEW: Linting stage - Checks code style and potential bugs
lint:
	@echo "Running golangci-lint..."
	golangci-lint run ./...

# NEW: Testing stage - Runs all unit tests
test:
	@echo "Running unit tests..."
	go test  ./...

build:
	go build -o $(BINARY_NAME) ./cmd/app/main.go

# NEW: Clean stage - Removes temporary binaries to keep the workspace clean
clean:
	rm -f $(BINARY_NAME) migrate admin

run:
	# Note: source doesn't work well inside Make. Better to use:
	# go run ./cmd/app/main.go (and let the OS handle env)
	go run ./cmd/app/main.go

env:
	set -a && source .env && set +a

generate-schema:
	go build -o migrate cmd/migrate/main.go && ./migrate generate-schema

migrate-up:
	go build -o migrate cmd/migrate/main.go && ./migrate up

migrate-down:
	go build -o migrate cmd/migrate/main.go && ./migrate down

create-admin:
	go build -o admin cmd/create-admin/main.go
	./admin $(EMAIL) $(PASSWORD)

# NEW: Docker Build - Packages the app for Docker Hub
docker-build:
	docker build -f ./Dockerfile.local -t $(DOCKER_USER)/$(BINARY_NAME):latest .
	docker tag $(DOCKER_USER)/$(BINARY_NAME):latest $(DOCKER_USER)/$(BINARY_NAME):$(shell git rev-parse --short HEAD)

docker-run:
	docker compose up --build
