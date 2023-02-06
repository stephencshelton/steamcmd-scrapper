# Use a minimal base image
FROM golang:alpine AS builder

# Set the working directory
WORKDIR /app

# Copy the go script
COPY entrypoint.go .

# Copy go.mod
COPY go.mod .

# Install the dependencies
RUN go get gopkg.in/yaml.v2

# Build the go script
RUN go build -o entrypoint entrypoint.go

# Use a minimal runtime image
FROM alpine

# Webhook url to push the scrapped information to
ENV WEBHOOK_URL=""
ENV WEBHOOK_TOKEN=""

# Set the working directory
WORKDIR /app

# Copy the yaml file
COPY applications.yaml .


# Copy the binary from the builder image
COPY --from=builder /app/entrypoint /app/entrypoint

# Set the entrypoint to the binary
ENTRYPOINT ["./entrypoint"]
