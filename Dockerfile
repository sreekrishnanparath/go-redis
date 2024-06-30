# Use the official Golang image as the base image
FROM golang:1.20 as builder

# Set the working directory in the container
WORKDIR /app

# Copy the package.json and package-lock.json files to the container
COPY go.mod .
COPY go.sum .

# Download the dependencies
RUN go mod download

# Copy the rest of the files to the container
COPY . .

# Build the application
RUN go build -o subscriber .

# Use the official Alpine image as the base image
FROM alpine:latest

# Set the working directory in the container
WORKDIR /app

# Copy the built application from the previous stage
COPY --from=builder /app/subscriber /app/subscriber

# Expose the port that the application will use
EXPOSE 8080

# Run the Redis server
RUN apk add --no-cache redis
CMD ["redis-server"]

# Run the application
CMD ["./subscriber"]