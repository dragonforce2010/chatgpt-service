FROM golang:1.19-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main ./cmd/http

FROM alpine:3.14

WORKDIR /app

COPY --from=build /app/main .

EXPOSE 8080

CMD ["./main"]
