FROM golang:1.23-alpine

ENV GOTOOLCHAIN=auto

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o quickwork ./cmd/api

EXPOSE 8080

CMD ["./quickwork"]
