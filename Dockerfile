FROM golang:latest AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/myapp cmd/api/*.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin/myapp .

EXPOSE 8080

CMD ["./myapp"]