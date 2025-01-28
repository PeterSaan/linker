FROM golang:1.23.5

WORKDIR /backend

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

COPY .env ./
