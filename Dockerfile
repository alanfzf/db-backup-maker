FROM alpine:latest

RUN apk add --no-cache \
    bash aws-cli mariadb-client mariadb-connector-c

# copy script
COPY --chmod=755 scripts/backup.sh /usr/local/bin/backup.sh

# start crond with log level 8 in foreground, output to stderr
CMD ["crond", "-f", "-d", "8"]
