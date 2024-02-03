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
COPY ./followpkg/ ./followpkg/
COPY ./helpers/ ./helpers/
COPY ./repos/ ./repos/
COPY ./userpkg/ ./userpkg/

RUN find . -name "*_test.go" -exec rm {} \;
RUN go build -o ./main ./cmd/app
RUN go clean -modcache
RUN rm -r ./auth ./cmd ./db ./followpkg ./helpers ./repos ./userpkg

COPY ./static ./static
COPY --from=scss_builder /app/static/index.css ./static/

COPY ./templates ./templates

EXPOSE 8080
CMD ["./main"]
