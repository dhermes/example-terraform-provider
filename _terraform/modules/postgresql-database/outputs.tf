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

output "admin_role" {
  value = module.admin_role.role_name
}

output "app_role" {
  value = module.app_role.role_name
}

output "db_name" {
  value = postgresql_database.db.name
}

output "schema" {
  value = postgresql_schema.application.name
}
