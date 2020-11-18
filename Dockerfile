FROM golang:alpine as builder
RUN apk update && apk add --no-cache git

WORKDIR /app

# For building Go Module required
ENV GOPROXY=direct
ENV GO111MODULE=on
ENV GOARCH=amd64
ENV GOOS=linux
ENV CGO_ENABLED=0

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN  go mod download

# Copy the go source
COPY . .

# Build
RUN  go build -a -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

# Expose port 9000 to the outside world
EXPOSE 9000

#Command to run the executable
CMD ["./main"]
