hello:
	echo "Hello"

build:
	go build -o app ./cmd/app/main.go

run:
	soure .env
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