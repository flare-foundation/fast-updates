# build executable
FROM golang:1.21 AS builder

WORKDIR /build

# Copy and download dependencies using go mod
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container
COPY . ./

# Build the applications
RUN go build -o /app/client ./main.go
RUN go build -o /app/keygen ./keygen/keygen.go

FROM debian:latest AS execution

WORKDIR /app

ENV PATH="/app/bin:${PATH}"
RUN mkdir bin

COPY --from=builder /app/client bin/client
COPY --from=builder /app/keygen bin/keygen

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["./bin/client", "--config", "config.toml"]
