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

resource "postgresql_grant" "grant_application_schema_to_admin" {
  role        = var.admin_role
  database    = var.db_name
  schema      = var.schema
  object_type = "schema"
  privileges  = ["USAGE", "CREATE"]
}

resource "postgresql_grant" "grant_application_schema_to_app" {
  role        = var.app_role
  database    = var.db_name
  schema      = var.schema
  object_type = "schema"
  privileges  = ["USAGE"]
}

resource "postgresql_grant" "app_table_grant" {
  database    = var.db_name
  role        = var.app_role
  schema      = var.schema
  object_type = "table"
  privileges  = ["SELECT", "DELETE", "INSERT", "UPDATE"]
}

resource "postgresql_grant" "app_seq_grant" {
  database    = var.db_name
  role        = var.app_role
  schema      = var.schema
  object_type = "sequence"
  privileges  = ["SELECT", "UPDATE"]
}

resource "postgresql_default_privileges" "app_table_grant" {
  database    = var.db_name
  role        = var.app_role
  schema      = var.schema
  owner       = var.admin_role
  object_type = "table"
  privileges  = ["SELECT", "DELETE", "INSERT", "UPDATE"]
}

resource "postgresql_default_privileges" "app_seq_grant" {
  database    = var.db_name
  role        = var.app_role
  schema      = var.schema
  owner       = var.admin_role
  object_type = "sequence"
  privileges  = ["SELECT", "UPDATE"]
}
