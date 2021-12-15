# example-terraform-provider

> Me just toying around with the problem space of writing a custom Terraform provider

## Development

```
$ make  # Or `make help`
Makefile for `example-terraform-provider` project

Usage:
   make vet                 Run `go vet` over source tree
   make shellcheck          Run `shellcheck` on all shell files in `./_bin/`
PostgreSQL-specific Targets:
   make start-postgres      Starts a PostgreSQL database running in a Docker container and set up users
   make stop-postgres       Stops the PostgreSQL database running in a Docker container
   make restart-postgres    Stops the PostgreSQL database (if running) and starts a fresh Docker container
   make require-postgres    Determine if PostgreSQL database is running; fail if not
   make psql                Connects to currently running PostgreSQL DB via `psql`
   make psql-superuser      Connects to currently running PostgreSQL DB via `psql` as superuser

```
