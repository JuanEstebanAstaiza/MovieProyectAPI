FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go build -o movieproyectapi

EXPOSE 8080

CMD ["./movieproyectapi"]
