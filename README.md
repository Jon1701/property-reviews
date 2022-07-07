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

- `make run` to run the Go application
