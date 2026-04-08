hello:
	echo "Hello"

build:
	go build -o app ./cmd/app/main.go

run:
	go run ./cmd/app/main.go

env:
	set -a && source .env && set +a

migrate-up:
	go build -o migrate cmd/migrate/main.go && ./migrate up

migrate-down:
	go build -o migrate cmd/migrate/main.go && ./migrate down

