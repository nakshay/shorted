
FROM golang:1.23.1-alpine AS builder

WORKDIR /app

COPY ../go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./bin/shorted


FROM alpine

WORKDIR /app

RUN mkdir configuration

COPY --from=builder /app/bin/shorted .
COPY ../configuration ./configuration

EXPOSE 8080

CMD ["./shorted"]
