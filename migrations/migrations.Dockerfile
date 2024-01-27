FROM golang:1.21.6-alpine

# Set the working directory inside the container
WORKDIR /app

# Install gcc and musl-dev for cgo
RUN apk add --no-cache gcc musl-dev

RUN CGO_CFLAGS="-D_LARGEFILE64_SOURCE" CGO_ENABLED=1 go install -tags='sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY ./migrations ./migrations

ENTRYPOINT [ "migrate", "-source", "file:///app/migrations", "-database", "sqlite3:///data/api.db", "up"]
