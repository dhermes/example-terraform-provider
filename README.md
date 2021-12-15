# example-terraform-provider

> Me just toying around with the problem space of writing a custom Terraform provider

## Development

```
$ make  # Or `make help`
Makefile for `example-terraform-provider` project

Usage:
   make serve                         Run the Books API application
   make seed-data                     Seed the database via the Books API
   make clean                         Forcefully remove all generated artifacts (e.g. Terraform state files)
   make vet                           Run `go vet` over source tree
   make shellcheck                    Run `shellcheck` on all shell files in `./_bin/`
Terraform-specific Targets:
   make install-terraform-provider    Install `terraform-provider-books` into Terraform plugins directory
   make apply-books-workspace         Apply the workspace that uses `terraform-provider-books`
   make start-postgres-container      Start PostgreSQL Docker containers
   make stop-postgres-container       Stop PostgreSQL Docker containers
   make initialize-database           Initialize the database, schema, roles and grants in the PostgreSQL instances
   make teardown-database             Teardown the database, schema, roles and grants in the PostgreSQL instances
PostgreSQL-specific Targets:
   make migrations-up                 Run PostgreSQL migrations for Books database
   make start-postgres                Starts a PostgreSQL database running in a Docker container and set up users
   make stop-postgres                 Stops the PostgreSQL database running in a Docker container
   make restart-postgres              Stops the PostgreSQL database (if running) and starts a fresh Docker container
   make clear-database                Deletes data from all existing tables
   make require-postgres              Determine if PostgreSQL database is running; fail if not
   make psql                          Connects to currently running PostgreSQL DB via `psql` as app user
   make psql-admin                    Connects to currently running PostgreSQL DB via `psql` as admin user
   make psql-superuser                Connects to currently running PostgreSQL DB via `psql` as superuser

```
