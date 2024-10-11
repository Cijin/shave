FROM node:18 AS tailwind-builder
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build-css

FROM golang:1.23 AS go-builder
WORKDIR /app
COPY . .

RUN go mod download
RUN go install github.com/a-h/templ/cmd/templ@v0.2.778

RUN templ generate
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags "-X shave/internal/version.Version=$(git describe --tags --always --dirty)" -o server ./cmd/server
RUN CGO_ENABLED=1 GOOS=linux go build -o migration ./cmd/migration

FROM golang:1.23.2-bookworm
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=tailwind-builder /app/public ./public
COPY --from=go-builder /app/server .
COPY --from=go-builder /app/migration .

EXPOSE 8080

ENTRYPOINT ["/bin/sh", "-c", "./migration && ./server"]
