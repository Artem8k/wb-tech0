
FROM golang:1.21.5-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./main main.go \
    && chmod +x ./main

ENTRYPOINT [ "./main" ]
