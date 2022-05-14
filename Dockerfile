FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY cmd ./cmd
COPY internal ./internal
COPY config.yaml ./

RUN go build cmd/main.go

EXPOSE 8000

CMD [ "./main" ]