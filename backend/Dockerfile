# GO Repo base repo
FROM golang:alpine as builder
RUN apk add git
RUN mkdir /app
WORKDIR /app

#COPY go.mod ./
COPY go.mod .
COPY go.sum .
# Download all the dependencies
RUN go mod download

COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -installsuffix cgo -o main .

# GO Repo base repo
FROM alpine:latest
RUN mkdir /app
WORKDIR /app/

ENV GIN_MODE=release
# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8080
EXPOSE 8080

# Run Executable
CMD ["./main"]