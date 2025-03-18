FROM golang:1.23.5

WORKDIR /backend

COPY . .

RUN go mod download
RUN bash setup_clinker.sh
