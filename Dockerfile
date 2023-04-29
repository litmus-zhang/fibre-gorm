FROM golang:1.16
LABEL authors="abdul"

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go mod download golang.org/x/crypto

COPY . .


CMD ["go", "run", "main.go"]