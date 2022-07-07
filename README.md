# Property Reviews App

## Status

Active

## Description

Full Stack Application which allows users to review residential properties.

## Requirements

- [Go v1.18](https://go.dev/)
- [GNU Make](https://www.gnu.org/software/make/)

## Environment Variables

| Variable                         | Description                                      |
| -------------------------------- | ------------------------------------------------ |
| `SERVER_PORT`                    | Backend server port                              |
| `POSTGRES_DOCKER_CONTAINER_NAME` | Container name for the Postgres Database service |
| `POSTGRES_ADMIN_USERNAME`        | Username for the Postgres administrator account  |
| `POSTGRES_ADMIN_PASSWORD`        | Password for the Postgres administrator account  |
| `POSTGRES_DB`                    | Default Postgres database                        |

## Commands

The following `make` commands are available:

- `make start-services` to start Docker services
- `make stop-services` to stop Docker services
- `make db-shell` to open a `bash` shell in the Postgres container
- `make psql` to open `psql` in the Postgres container
- `make run` to run the Go application
- `make initialize-db` to initialize the database (create users, tables, etc)
