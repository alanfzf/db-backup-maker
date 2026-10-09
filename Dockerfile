FROM golang:1.26-alpine AS builder

# Do layer caching
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

# Build the actual app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux \
    go build \
    -trimpath -ldflags="-s -w" \
    -o /out/db-backup-maker \
    .

FROM alpine:3.24 AS runtime

RUN apk add --no-cache \
    mariadb-client \
    supercronic

# Copy compiled Go application
COPY --from=builder /out/db-backup-maker /usr/local/bin/db-backup-maker

CMD ["/usr/bin/supercronic", "/etc/crontab"]
