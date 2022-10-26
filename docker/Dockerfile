FROM golang:1.18-alpine as builder
WORKDIR /tmp/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out .

FROM golang:1.18

RUN go install github.com/githubnemo/CompileDaemon@latest

WORKDIR /app
COPY --from=builder /tmp/app/out /app

ENTRYPOINT CompileDaemon --command=./fintech_test_project
