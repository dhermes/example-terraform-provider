# example-terraform-provider

> Me just toying around with the problem space of writing a custom Terraform
> provider

## Prerequisites

Bring up the database (and ensure migrations run, etc.)

```bash
make restart-postgres
# It's a bit janky, there is a race condition with grants causing occasional
# errors "tuple concurrently updated"
```

Run the Books API server

```bash
make serve
```

Clear / seed the database with some data

```bash
make clear-database
make seed-data
```

Install the provider into `~/.terraform.d/plugins` (or a different Terraform
plugins directory if configured):

```
make install-terraform-provider
```

## Demo

First setup / cleanup the workspace:

```bash
# Or: `make apply-books-workspace`

cd _terraform/workspaces/books/
# Technically should also remove the 'George R.R. Martin' row from the
# database if it's still present there; can just do this via
# `make clear-database && make seed-data`.
rm --force .terraform.lock.hcl terraform.tfstate terraform.tfstate.backup
rm --force --recursive .terraform/
terraform init
```

After doing this, our provider should have been pulled from the local
Terraform plugins directory:

```
$ tree .terraform/
.terraform/
└── providers
    └── tf-registry.invalid
        └── dhermes
            └── books
                └── 0.0.1
                    └── linux_amd64 -> /home/dhermes/.terraform.d/plugins/tf-registry.invalid/dhermes/books/0.0.1/linux_amd64

6 directories, 0 files
```

then apply the workspace:

```bash
terraform apply
```

Then check out the newly created resources:

```
$ terraform apply
...
  Enter a value: yes

books_api_author.grr_martin: Creating...
books_api_author.grr_martin: Creation complete after 0s [id=320a97c7-7a44-4a73-8b5c-a3d4c701ed67]

Apply complete! Resources: 1 added, 0 changed, 0 destroyed.
$
$
$ cat terraform.tfstate | jq '.resources'
[
  {
    "mode": "managed",
    "type": "books_api_author",
    "name": "grr_martin",
    "provider": "provider[\"tf-registry.invalid/dhermes/books\"].local",
    "instances": [
      {
        "schema_version": 0,
        "attributes": {
          "first_name": "George R.R.",
          "id": "320a97c7-7a44-4a73-8b5c-a3d4c701ed67",
          "last_name": "Martin"
        },
        "sensitive_attributes": [],
        "private": "bnVsbA=="
      }
    ]
  }
]
$
$
$ psql --dbname 'postgres://books_app:testpassword_app@127.0.0.1:22411/books' --command "SELECT * FROM authors WHERE last_name = 'Martin'"
                  id                  | first_name  | last_name
--------------------------------------+-------------+-----------
 320a97c7-7a44-4a73-8b5c-a3d4c701ed67 | George R.R. | Martin
(1 row)

$
```

To make sure other state transitions are in good shape, apply it again:

```bash
terraform apply
terraform apply
# ...
```

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
   make destroy-books-workspace       Destroy the workspace that uses `terraform-provider-books`
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
