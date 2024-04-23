
FROM golang:1.19-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/weather-api cmd/main.go
FROM alpine:latest

WORKDIR /app

COPY static/ .

COPY --from=build /app/weather-api .




EXPOSE 8080


CMD ["./weather-api"]