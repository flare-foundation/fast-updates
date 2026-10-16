#checkov:skip=CKV_DOCKER_2: Healthcheck instruction not required
# build executable
FROM golang:1.27.2-trixie@sha256:47667934874913a449b96faa7ef4777fa285d9ec29126ee80ed61c45b7cccf0e AS builder

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
