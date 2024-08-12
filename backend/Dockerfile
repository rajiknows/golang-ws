# Build Stage
FROM golang:1.22-alpine AS build
WORKDIR /app
# Copy Go mod files and download dependencies
COPY go.* ./
RUN go mod download
# Install Goose
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
# Copy the rest of the application code
COPY . .
# Build the Go application
RUN go build -o main main.go

# Runtime Stage
FROM golang:1.22-alpine
WORKDIR /app

# Install bash
RUN apk add --no-cache bash

# Copy the built binary and Goose from the build stage
COPY --from=build /app/ /app/
COPY --from=build /go/bin/goose /usr/local/bin/goose

# Copy the wait-for-it script and make it executable
COPY wait-for-it.sh /usr/local/bin/wait-for-it.sh
RUN chmod +x /usr/local/bin/wait-for-it.sh


# Copy the start script and make it executable
COPY start.sh /app/start.sh
RUN chmod +x /app/start.sh

# Expose the port
EXPOSE 8080

# Set the entry point to the start script
CMD ["/bin/bash", "/app/start.sh"]
