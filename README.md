# MARIA-BACKUP-MAKER

Containerized MariaDB backup job that runs on cron, dumps all databases, compresses the dump, and uploads it to S3.

## Environment variables

Create a `.env` file in this directory (or copy from `.env.example`):

```bash
cp .env.example .env
```

| Variable | Required | Default | Description |
|---|---|---|---|
| `AWS_REGION` | No | `us-east-1` | AWS region used by AWS CLI |
| `AWS_ACCESS_KEY_ID` | Yes | - | AWS access key |
| `AWS_SECRET_ACCESS_KEY` | Yes | - | AWS secret key |
| `AWS_BUCKET` | Yes | - | Full S3 destination path, for example `s3://my-backups/mariadb` |
| `DB_HOST` | No | `host.docker.internal` | MariaDB hostname reachable from container |
| `DB_PORT` | No | `3306` | MariaDB port |
| `DB_USER` | Yes | - | MariaDB user with dump permissions |
| `DB_PASS` | Yes | - | MariaDB user password |

## Setup

1. Copy env template and fill real values:
   ```bash
   cp .env.example .env
   ```
2. Build and start:
   ```bash
   docker compose up -d --build
   ```

## Cron schedule

The container reads its cron schedule from `config/crontab`, which is bind-mounted into `/etc/crontabs/root`.

Default `config/crontab`:

```cron
# ┌ minute (0–59)
# │ ┌ hour (0–23)
# │ │ ┌ day of month (1–31)
# │ │ │ ┌ month (1–12)
# │ │ │ │ ┌ day of week (0–7, Sun=0 or 7)
# │ │ │ │ │
# * * * * * command_to_execute
# 00:00 CST
00 06 * * * backup.sh
# 13:15 CST
15 19 * * * backup.sh
# 18:15 CST
15 00 * * * backup.sh
```

### Multiple schedules

Add multiple cron lines to the same `config/crontab`, one per run. For example:

```cron
00 06 * * * backup.sh
15 19 * * * backup.sh
15 00 * * * backup.sh
```

After editing the file, reload with:

```bash
docker compose restart
```
## Run an immediate backup (manual)

```bash
docker compose run --rm maria-backup-maker backup.sh
```
