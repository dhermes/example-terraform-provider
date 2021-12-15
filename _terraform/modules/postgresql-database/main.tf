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

module "admin_role" {
  source = "../postgresql-role"

  username = "${var.db_name}_admin"
  password = var.admin_password
}

resource "postgresql_database" "db" {
  name              = var.db_name
  owner             = module.admin_role.role_name
  template          = "template0"
  encoding          = "UTF8"
  lc_collate        = "en_US.UTF-8"
  lc_ctype          = "en_US.UTF-8"
  connection_limit  = -1
  allow_connections = true
}

module "app_role" {
  source = "../postgresql-role"

  username = "${var.db_name}_app"
  password = var.app_password
}

resource "postgresql_extension" "pgcrypto" {
  name     = "pgcrypto"
  schema   = "public"
  database = postgresql_database.db.name
}
