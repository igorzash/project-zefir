FROM alpine:3.19.1 as scss_builder
RUN apk add --no-cache nodejs npm
WORKDIR /app

RUN npm install -g sass

COPY ./styles ./styles
RUN sass  --style compressed ./styles/index.scss ./static/index.css

FROM alpine:3.19.1 as runtime
RUN apk add --no-cache go gcc musl-dev
WORKDIR /app

COPY go.mod go.sum ./

ENV CGO_ENABLED=1
RUN go mod download

COPY ./app/ ./app/
COPY ./auth/ ./auth/
COPY ./cmd/ ./cmd/
COPY ./controllers/ ./controllers/
COPY ./db/ ./db/
COPY ./entities/ ./entities/
COPY ./helpers/ ./helpers/
COPY ./services/ ./services/

RUN find . -name "*_test.go" -exec rm {} \;
RUN go build -o ./main ./cmd/app
RUN go clean -modcache
RUN rm -r ./app ./auth ./cmd ./controllers ./db ./entities ./helpers ./services

COPY ./static ./static
COPY --from=scss_builder /app/static/index.css ./static/

COPY ./templates ./templates

EXPOSE 8080
CMD ["./main"]
