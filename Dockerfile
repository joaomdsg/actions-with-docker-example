# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.19-buster AS builder

# Create and change to the app directory.
WORKDIR /app

# Copy the Go module files to the container.
COPY go.mod go.sum ./

# Download the dependencies.
RUN go mod download

# Copy the rest of the application code into the container.
COPY . .

# Build the application inside the container.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Use a minimal base image to package the app binary.
# https://hub.docker.com/_/scratch
FROM scratch

# Copy the binary from the builder image.
COPY --from=builder /app/app /

# Run the app binary.
CMD ["/app"]
