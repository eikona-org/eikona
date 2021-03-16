FROM golang:alpine as debug

RUN apk update && apk upgrade && \
    apk add --no-cache git \
        dpkg \
        gcc \
        git \
        musl-dev

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
# Download all the dependencies
RUN go mod download

RUN go get github.com/go-delve/delve/cmd/dlv

COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -installsuffix cgo -o main .

WORKDIR /

COPY ./dlv.sh /
RUN chmod +x /dlv.sh
ENTRYPOINT ["/dlv.sh"]

#########################################################################
# GO Repo base repo
FROM golang:alpine as builder
RUN apk add git
RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
# Download all the dependencies
RUN go mod download

COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -installsuffix cgo -o main .

#########################################################################
# GO Repo base repo
FROM alpine:latest as prod
RUN apk --no-cache add ca-certificates curl
RUN mkdir /app
WORKDIR /app/

ENV GIN_MODE=release
# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8080
EXPOSE 8080

# Run Executable
CMD ["./main"]