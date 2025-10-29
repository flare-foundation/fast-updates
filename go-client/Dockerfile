# build executable
FROM golang:1.21-bookworm@sha256:c6a5b9308b3f3095e8fde83c8bf4d68bd101fce606c1a0a1394522542509dda9 AS builder

WORKDIR /build

# Copy and download dependencies using go mod
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container
COPY . ./

# Build the applications
RUN go build -o /app/client ./main.go
RUN go build -o /app/keygen ./keygen/keygen.go

FROM debian:trixie@sha256:72547dd722cd005a8c2aa2079af9ca0ee93aad8e589689135feaed60b0a8c08d AS execution

WORKDIR /app

ENV PATH="/app/bin:${PATH}"
RUN mkdir bin

COPY --from=builder /app/client bin/client
COPY --from=builder /app/keygen bin/keygen

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["./bin/client", "--config", "config.toml"]
