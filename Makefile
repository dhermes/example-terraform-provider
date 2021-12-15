# Copyright 2021 Danny Hermes
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

.PHONY: help
help:
	@echo 'Makefile for `example-terraform-provider` project'
	@echo ''
	@echo 'Usage:'
	@echo '   make serve                       Run the Books API application'
	@echo '   make seed-data                   Seed the database via the Books API'
	@echo '   make clean                       Forcefully remove all generated artifacts (e.g. Terraform state files)'
	@echo '   make vet                         Run `go vet` over source tree'
	@echo '   make shellcheck                  Run `shellcheck` on all shell files in `./_bin/`'
	@echo 'Terraform-specific Targets:'
	@echo '   make start-postgres-container    Start PostgreSQL Docker containers.'
	@echo '   make stop-postgres-container     Stop PostgreSQL Docker containers.'
	@echo '   make initialize-database         Initialize the database, schema, roles and grants in the PostgreSQL instances'
	@echo '   make teardown-database           Teardown the database, schema, roles and grants in the PostgreSQL instances'
	@echo 'PostgreSQL-specific Targets:'
	@echo '   make migrations-up               Run PostgreSQL migrations for Books database'
	@echo '   make start-postgres              Starts a PostgreSQL database running in a Docker container and set up users'
	@echo '   make stop-postgres               Stops the PostgreSQL database running in a Docker container'
	@echo '   make restart-postgres            Stops the PostgreSQL database (if running) and starts a fresh Docker container'
	@echo '   make clear-database              Deletes data from all existing tables'
	@echo '   make require-postgres            Determine if PostgreSQL database is running; fail if not'
	@echo '   make psql                        Connects to currently running PostgreSQL DB via `psql` as app user'
	@echo '   make psql-admin                  Connects to currently running PostgreSQL DB via `psql` as admin user'
	@echo '   make psql-superuser              Connects to currently running PostgreSQL DB via `psql` as superuser'
	@echo ''

################################################################################
# Meta-variables
################################################################################
SHELLCHECK_PRESENT := $(shell command -v shellcheck 2> /dev/null)

################################################################################
# Environment variable defaults
################################################################################
DB_HOST ?= 127.0.0.1
DB_PORT ?= 22411

DB_SUPERUSER_NAME ?= superuser_db
DB_SUPERUSER_USER ?= superuser
DB_SUPERUSER_PASSWORD ?= testpassword_superuser

DB_NAME ?= books
DB_APP_USER ?= books_app
DB_APP_PASSWORD ?= testpassword_app
DB_ADMIN_USER ?= books_admin
DB_ADMIN_PASSWORD ?= testpassword_admin

DB_SCHEMA ?= books
DB_METADATA_TABLE ?= books_migrations

# NOTE: This assumes the `DB_*_PASSWORD` values do not need to be URL encoded.
SUPERUSER_DSN ?= postgres://$(DB_SUPERUSER_USER):$(DB_SUPERUSER_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_SUPERUSER_NAME)
APP_DSN ?= postgres://$(DB_APP_USER):$(DB_APP_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)
ADMIN_DSN ?= postgres://$(DB_ADMIN_USER):$(DB_ADMIN_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)

SERVER_BIND_PORT ?= 7534
SERVER_BIND_ADDR ?= :$(SERVER_BIND_PORT)
SEED_BOOKS_ADDR ?= http://localhost:$(SERVER_BIND_PORT)

################################################################################
# Generic Targets
################################################################################

.PHONY: serve
serve:
	go run ./cmd/server/ --addr $(SERVER_BIND_ADDR) --dsn $(APP_DSN)

.PHONY: seed-data
seed-data:
	BOOKS_ADDR=$(SEED_BOOKS_ADDR) ./_bin/seed_data.sh

.PHONY: clean
clean:
	rm --force \
	  _terraform/workspaces/database/.terraform.lock.hcl \
	  _terraform/workspaces/database/terraform.tfstate \
	  _terraform/workspaces/database/terraform.tfstate.backup \
	  _terraform/workspaces/docker/.terraform.lock.hcl \
	  _terraform/workspaces/docker/terraform.tfstate \
	  _terraform/workspaces/docker/terraform.tfstate.backup
	rm --force --recursive \
	  _terraform/workspaces/database/.terraform/ \
	  _terraform/workspaces/docker/.terraform/
	docker rm --force \
	  dev-postgres-books
	docker network rm dev-network-books || true

.PHONY: vet
vet:
	go vet ./...

.PHONY: shellcheck
shellcheck: _require-shellcheck
	shellcheck --exclude SC1090 ./_bin/*.sh

################################################################################
# Terraform-specific Targets
################################################################################

.PHONY: start-postgres-container
start-postgres-container:
	@cd _terraform/workspaces/docker/ && \
	  terraform init && \
	  terraform apply --auto-approve

.PHONY: stop-postgres-container
stop-postgres-container:
	@cd _terraform/workspaces/docker/ && \
	  terraform init && \
	  terraform apply --destroy --auto-approve

.PHONY: initialize-database
initialize-database:
	@cd _terraform/workspaces/database/ && \
	  terraform init && \
	  terraform apply --auto-approve

.PHONY: teardown-database
teardown-database:
	@cd _terraform/workspaces/database/ && \
	  terraform init && \
	  terraform apply --destroy --auto-approve

################################################################################
# PostgreSQL
################################################################################

.PHONY: migrations-up
migrations-up:
	@PGPASSWORD=$(DB_ADMIN_PASSWORD) go run ./cmd/migrations-up/ \
	  --dev \
	  --metadata-table $(DB_METADATA_TABLE) \
	  postgres \
	  --dbname $(DB_NAME) \
	  --driver-name pgx \
	  --port $(DB_PORT) \
	  --schema $(DB_SCHEMA) \
	  --username $(DB_ADMIN_USER) \
	  up

.PHONY: start-postgres
start-postgres: start-postgres-container require-postgres initialize-database migrations-up

.PHONY: stop-postgres
stop-postgres: teardown-database stop-postgres-container

.PHONY: restart-postgres
restart-postgres: stop-postgres start-postgres

.PHONY: clear-database
clear-database: require-postgres
	@DB_FULL_DSN=$(ADMIN_DSN) ./_bin/clear_database.sh

.PHONY: require-postgres
require-postgres:
	@DB_HOST=$(DB_HOST) \
	  DB_PORT=$(DB_PORT) \
	  DB_FULL_DSN=$(ADMIN_DSN) \
	  ./_bin/require_postgres.sh

.PHONY: psql
psql: require-postgres
	@echo "Running psql against port $(DB_PORT)"
	psql "$(APP_DSN)"

.PHONY: psql-admin
psql-admin: require-postgres
	@echo "Running psql against port $(DB_PORT)"
	psql "$(ADMIN_DSN)"

.PHONY: psql-superuser
psql-superuser: require-postgres
	@echo "Running psql against port $(DB_PORT)"
	psql "$(SUPERUSER_DSN)"

################################################################################
# Internal / Doctor Targets
################################################################################

.PHONY: _require-shellcheck
_require-shellcheck:
ifndef SHELLCHECK_PRESENT
	$(error 'shellcheck is not installed, it can be installed via "apt-get install shellcheck" or "brew install shellcheck".')
endif
