#checkov:skip=CKV_DOCKER_2: Healthcheck instruction not required
# build executable
FROM golang:1.25.14-trixie@sha256:2c4c60ef415fbfa5e90300722293bef36c5e63fae17570ce18f580af933dbd73 AS builder

WORKDIR /build

# Copy and download dependencies using go mod
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container
COPY . ./

# Build the applications
RUN go build -o /app/client ./main.go && \
    go build -o /app/keygen ./keygen/keygen.go

FROM debian:trixie@sha256:72547dd722cd005a8c2aa2079af9ca0ee93aad8e589689135feaed60b0a8c08d AS execution

WORKDIR /app

ENV PATH="/app/bin:${PATH}"
RUN mkdir bin

COPY --from=builder /app/client bin/client
COPY --from=builder /app/keygen bin/keygen

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

RUN groupadd -g 10001 app && useradd -u 10001 -g app -s /sbin/nologin app && chown -R app:app /app

USER app

CMD ["./bin/client", "--config", "config.toml"]
