# MARIA-BACKUP-MAKER

Containerized MariaDB backup job that dumps all databases, compresses the dump, and uploads it to S3.

## Environment variables

Create a `.env` file in this directory (or copy from `.env.example`):

```bash
cp .env.example .env
```

## Cron Schedule

Add multiple cron lines to the same `config/crontab`, one per run. For example:

```bash
# ┌ minute (0–59)
# │ ┌ hour (0–23)
# │ │ ┌ day of month (1–31)
# │ │ │ ┌ month (1–12)
# │ │ │ │ ┌ day of week (0–7, Sun=0 or 7)
# │ │ │ │ │
# * * * * * command_to_execute

# 00:35 CST
35 06 * * * db-backup-maker
# 13:35 CST
35 19 * * * db-backup-maker
# 18:35 CST
35 00 * * * db-backup-maker
```

## Run an immediate backup (manual)

```bash
docker compose run --rm maria-backup-maker db-backup-maker
```
