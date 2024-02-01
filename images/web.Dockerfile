FROM node:21-alpine3.18 as scss_builder
WORKDIR /app

RUN npm install -g sass

COPY ./styles ./styles
RUN sass  --style compressed ./styles/index.scss ./static/index.css

FROM golang:1.21.6-alpine as runtime
WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY ./auth/ ./auth/
COPY ./cmd/ ./cmd/
COPY ./db/ ./db/
COPY ./followpkg/ ./followpkg/
COPY ./helpers/ ./helpers/
COPY ./repos/ ./repos/
COPY ./userpkg/ ./userpkg/

RUN find . -name "*_test.go" -exec rm {} \; && \
    CGO_ENABLED=1 go build -o ./main ./cmd/app && \
    go clean -modcache && \
    rm -r ./auth ./cmd ./db ./followpkg ./helpers ./repos ./userpkg

COPY ./static ./static
COPY --from=scss_builder /app/static/index.css ./static/

COPY ./templates ./templates

EXPOSE 8080
CMD ["./main"]
