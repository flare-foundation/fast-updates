# build executable
FROM golang:1.25-trixie@sha256:c699a45b4ed445f5a618233d1910f7c889dd52b831594047eff8e6a52794c814 AS builder

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
