# Use an official Golang runtime as the base image
FROM golang:1.19-alpine AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o /app/weather-api cmd/main.go
# Start a new stage from scratch
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

COPY static/ .

# Copy the Pre-built binary file from the previous stage
COPY --from=build /app/weather-api .




# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable

CMD ["./weather-api"]