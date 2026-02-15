FROM golang:1.24 AS base

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main main.go

CMD [ "./main" ]