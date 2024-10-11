# CSS
FROM node:18 AS tailwind-builder
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build-css

# Server
FROM golang:1.23 AS go-builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/a-h/templ/cmd/templ@v0.2.778
COPY . .

RUN templ generate

# Build server
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X paper-chase/internal/version.Version=$(git describe --tags --always --dirty)" -o server ./cmd/server

# Build migrations
RUN CGO_ENABLED=0 GOOS=linux go build -o migration ./cmd/migration

# Final Image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app

COPY --from=tailwind-builder /app/public ./public
COPY --from=go-builder /app/server .

COPY --from=go-builder /app/migration .

EXPOSE 8080

ENTRYPOINT ["/bin/sh", "-c", "./migration && ./server"]
