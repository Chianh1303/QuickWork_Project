FROM golang:alpine

ENV GOTOOLCHAIN=auto
ENV PORT=8080

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o quickwork ./cmd/api

EXPOSE 8080

CMD ["./quickwork"]
