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

module "books_database" {
  source = "../../modules/postgresql-database"

  db_name        = "books"
  admin_password = "testpassword_admin"
  app_password   = "testpassword_app"
  schema         = "books"

  providers = {
    postgresql = postgresql.docker
  }
}

module "books_grants" {
  source = "../../modules/postgresql-grants"

  schema     = module.books_database.schema
  db_name    = module.books_database.db_name
  admin_role = module.books_database.admin_role
  app_role   = module.books_database.app_role

  providers = {
    postgresql = postgresql.docker
  }
}
