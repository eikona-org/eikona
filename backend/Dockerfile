FROM golang:alpine as debug

RUN apk add git curl

WORKDIR /app

RUN go get github.com/go-delve/delve/cmd/dlv

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

CMD air -c .air.toml

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

# Build the docs
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN swag init

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -installsuffix cgo -o main .

#########################################################################
# GO Repo base repo
FROM alpine:latest as prod
RUN mkdir /app
WORKDIR /app/

ENV GIN_MODE=release
# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8080
EXPOSE 8080

# Run Executable
CMD ["./main"]
