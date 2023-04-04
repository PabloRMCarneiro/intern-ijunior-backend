FROM golang:1.20.2

WORKDIR /usr/src/app

RUN go install github.com/githubnemo/CompileDaemon

COPY . .

RUN go mod tidy