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
books_api_book.sirens: Creating...
books_api_author.grr_martin: Creation complete after 0s [id=04a76223-6d29-43d7-bcae-744773efd964]
books_api_book.song_fire_ice5: Creating...
books_api_book.song_fire_ice2: Creating...
books_api_book.song_fire_ice3: Creating...
books_api_book.song_fire_ice1: Creating...
books_api_book.song_fire_ice4: Creating...
books_api_book.song_fire_ice4: Creation complete after 0s [id=b368334e-4c05-4d55-ab7a-0c8fd1afa1c5]
books_api_book.sirens: Creation complete after 0s [id=d1f0f231-554d-4e2a-8c53-a007ad67763b]
books_api_book.song_fire_ice3: Creation complete after 0s [id=af2a723a-1fce-415d-8024-2ead084493ed]
books_api_book.song_fire_ice1: Creation complete after 0s [id=1ff1b78c-37b1-45b3-8ca8-93219edf2ea4]
books_api_book.song_fire_ice5: Creation complete after 0s [id=bc2fab4f-2d30-495a-8c9c-7324aaed93b1]
books_api_book.song_fire_ice2: Creation complete after 0s [id=9001ae4c-6eba-40ef-bdb1-1d95961201e3]

Apply complete! Resources: 7 added, 0 changed, 0 destroyed.
$
$
$ jq '.resources | length' terraform.tfstate
8
$
$
$ jq '.resources[0]' terraform.tfstate
{
  "mode": "data",
  "type": "books_api_author",
  "name": "vonnegut",
  "provider": "provider[\"tf-registry.invalid/dhermes/books\"].local",
  "instances": [
    {
      "schema_version": 0,
      "attributes": {
        "book_count": 0,
        "first_name": "Kurt",
        "id": "f4c5a610-753a-4940-8314-c5dff54477af",
        "last_name": "Vonnegut"
      },
      "sensitive_attributes": []
    }
  ]
}
$
$
$ jq '.resources[1]' terraform.tfstate
{
  "mode": "managed",
  "type": "books_api_author",
  "name": "grr_martin",
  "provider": "provider[\"tf-registry.invalid/dhermes/books\"].local",
  "instances": [
    {
      "schema_version": 0,
      "attributes": {
        "book_count": 0,
        "first_name": "George R.R.",
        "id": "04a76223-6d29-43d7-bcae-744773efd964",
        "last_name": "Martin"
      },
      "sensitive_attributes": [],
      "private": "bnVsbA=="
    }
  ]
}
$
$
$ jq '.resources[2]' terraform.tfstate
{
  "mode": "managed",
  "type": "books_api_book",
  "name": "sirens",
  "provider": "provider[\"tf-registry.invalid/dhermes/books\"].local",
  "instances": [
    {
      "schema_version": 0,
      "attributes": {
        "author_id": "f4c5a610-753a-4940-8314-c5dff54477af",
        "id": "d1f0f231-554d-4e2a-8c53-a007ad67763b",
        "publish_date": "1959-01-01",
        "title": "The Sirens of Titan"
      },
      "sensitive_attributes": [],
      "private": "bnVsbA==",
      "dependencies": [
        "data.books_api_author.vonnegut"
      ]
    }
  ]
}
$
$
$ psql --dbname 'postgres://books_app:testpassword_app@127.0.0.1:22411/books' --command "SELECT * FROM authors WHERE last_name IN ('Martin', 'Vonnegut')"
                  id                  | first_name  | last_name
--------------------------------------+-------------+-----------
 f4c5a610-753a-4940-8314-c5dff54477af | Kurt        | Vonnegut
 04a76223-6d29-43d7-bcae-744773efd964 | George R.R. | Martin
(2 rows)

$
$
$ psql --dbname 'postgres://books_app:testpassword_app@127.0.0.1:22411/books' --command "SELECT * FROM books WHERE author_id IN ('f4c5a610-753a-4940-8314-c5dff54477af', '04a76223-6d29-43d7-bcae-744773efd964')"
                  id                  |              author_id               |        title         |      publish_date
--------------------------------------+--------------------------------------+----------------------+------------------------
 b368334e-4c05-4d55-ab7a-0c8fd1afa1c5 | 04a76223-6d29-43d7-bcae-744773efd964 | A Feast for Crows    | 2005-10-17 00:00:00+00
 d1f0f231-554d-4e2a-8c53-a007ad67763b | f4c5a610-753a-4940-8314-c5dff54477af | The Sirens of Titan  | 1959-01-01 00:00:00+00
 1ff1b78c-37b1-45b3-8ca8-93219edf2ea4 | 04a76223-6d29-43d7-bcae-744773efd964 | A Game of Thrones    | 1996-08-01 00:00:00+00
 af2a723a-1fce-415d-8024-2ead084493ed | 04a76223-6d29-43d7-bcae-744773efd964 | A Storm of Swords    | 2000-08-08 00:00:00+00
 9001ae4c-6eba-40ef-bdb1-1d95961201e3 | 04a76223-6d29-43d7-bcae-744773efd964 | A Clash of Kings     | 1998-11-16 00:00:00+00
 bc2fab4f-2d30-495a-8c9c-7324aaed93b1 | 04a76223-6d29-43d7-bcae-744773efd964 | A Dance with Dragons | 2011-07-12 00:00:00+00
(6 rows)

$
```

To make sure other state transitions are in good shape, apply it again:

```
$ terraform apply
...
Note: Objects have changed outside of Terraform

Terraform detected the following changes made outside of Terraform since the last "terraform apply":

  # books_api_author.grr_martin has changed
  ~ resource "books_api_author" "grr_martin" {
      ~ book_count = 0 -> 5
        id         = "04a76223-6d29-43d7-bcae-744773efd964"
        # (2 unchanged attributes hidden)
    }


Unless you have made equivalent changes to your configuration, or ignored the relevant attributes using ignore_changes, the following plan may include actions to undo or respond to these changes.

─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────

No changes. Your infrastructure matches the configuration.

Your configuration already matches the changes detected above. If you'd like to update the Terraform state to match, create and apply a refresh-only plan:
  terraform apply -refresh-only

Apply complete! Resources: 0 added, 0 changed, 0 destroyed.
$
$
$ terraform apply
books_api_author.grr_martin: Refreshing state... [id=04a76223-6d29-43d7-bcae-744773efd964]
books_api_book.sirens: Refreshing state... [id=d1f0f231-554d-4e2a-8c53-a007ad67763b]
books_api_book.song_fire_ice2: Refreshing state... [id=9001ae4c-6eba-40ef-bdb1-1d95961201e3]
books_api_book.song_fire_ice1: Refreshing state... [id=1ff1b78c-37b1-45b3-8ca8-93219edf2ea4]
books_api_book.song_fire_ice4: Refreshing state... [id=b368334e-4c05-4d55-ab7a-0c8fd1afa1c5]
books_api_book.song_fire_ice5: Refreshing state... [id=bc2fab4f-2d30-495a-8c9c-7324aaed93b1]
books_api_book.song_fire_ice3: Refreshing state... [id=af2a723a-1fce-415d-8024-2ead084493ed]

No changes. Your infrastructure matches the configuration.

Terraform has compared your real infrastructure against your configuration and found no differences, so no changes are needed.

Apply complete! Resources: 0 added, 0 changed, 0 destroyed.
```

(Homework: make a small edit to `_terraform/workspaces/books/vonnegut.tf`
to make sure updates work.)

Now, test out deletes for the resources:

```bash
# Or: `make destroy-books-workspace`
terraform apply --destroy
```

Running this with the examples above:

```
$ terraform apply --destroy
...
  Enter a value: yes

books_api_book.sirens: Destroying... [id=d1f0f231-554d-4e2a-8c53-a007ad67763b]
books_api_book.song_fire_ice1: Destroying... [id=1ff1b78c-37b1-45b3-8ca8-93219edf2ea4]
books_api_book.song_fire_ice2: Destroying... [id=9001ae4c-6eba-40ef-bdb1-1d95961201e3]
books_api_book.song_fire_ice5: Destroying... [id=bc2fab4f-2d30-495a-8c9c-7324aaed93b1]
books_api_book.song_fire_ice3: Destroying... [id=af2a723a-1fce-415d-8024-2ead084493ed]
books_api_book.song_fire_ice4: Destroying... [id=b368334e-4c05-4d55-ab7a-0c8fd1afa1c5]
books_api_book.song_fire_ice2: Destruction complete after 0s
books_api_book.song_fire_ice4: Destruction complete after 0s
books_api_book.song_fire_ice1: Destruction complete after 0s
books_api_book.sirens: Destruction complete after 0s
books_api_book.song_fire_ice5: Destruction complete after 0s
books_api_book.song_fire_ice3: Destruction complete after 0s
books_api_author.grr_martin: Destroying... [id=04a76223-6d29-43d7-bcae-744773efd964]
books_api_author.grr_martin: Destruction complete after 0s

Apply complete! Resources: 0 added, 0 changed, 7 destroyed.
$
$
$ psql --dbname 'postgres://books_app:testpassword_app@127.0.0.1:22411/books' --command "SELECT * FROM authors WHERE last_name IN ('Martin', 'Vonnegut')"                  id                  | first_name | last_name
--------------------------------------+------------+-----------
 f4c5a610-753a-4940-8314-c5dff54477af | Kurt       | Vonnegut
(1 row)

$
$
$ psql --dbname 'postgres://books_app:testpassword_app@127.0.0.1:22411/books' --command "SELECT * FROM books WHERE author_id IN ('f4c5a610-753a-4940-8314-c5dff54477af', '04a76223-6d29-43d7-bcae-744773efd964')" id | author_id | title | publish_date
----+-----------+-------+--------------
(0 rows)

$
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
