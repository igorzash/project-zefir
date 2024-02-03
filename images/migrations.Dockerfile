FROM alpine:3.19.1
RUN apk add --no-cache go gcc musl-dev
WORKDIR /app

ENV CGO_ENABLED=1
RUN go install -tags='sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY ./migrations ./migrations

ENTRYPOINT [ "migrate", "-source", "file:///app/migrations", "-database", "sqlite3:///data/web.db", "up"]
