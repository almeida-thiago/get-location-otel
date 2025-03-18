FROM golang:1.23

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/weather cmd/main.go

EXPOSE 8080

ENTRYPOINT ["/app/weather"]