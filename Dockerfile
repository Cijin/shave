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
# CGO needs to be enabled as go-libsql required it
RUN GOOS=linux go build -ldflags "-X shave/internal/version.Version=$(git describe --tags --always --dirty)" -o server ./cmd/server
RUN GOOS=linux go build -o migration ./cmd/migration

# Final Image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app

COPY --from=tailwind-builder /app/public ./public
COPY --from=go-builder /app/server .

COPY --from=go-builder /app/migration .

EXPOSE 8080

ENTRYPOINT ["/bin/sh", "-c", "./migration && ./server"]
