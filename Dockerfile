# ---------- Build ----------
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o app ./cmd/app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o migrate ./cmd/migrate
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o create-admin ./cmd/create-admin


# ---------- Runtime ----------
FROM gcr.io/distroless/static-debian12

WORKDIR /

COPY --from=builder /app/app /app
COPY --from=builder /app/migrate /migrate
COPY --from=builder /app/create-admin /create-admin

USER nonroot:nonroot

ENTRYPOINT ["/app"]