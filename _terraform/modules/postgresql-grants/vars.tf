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

variable "schema" {
  type        = string
  description = "The name of the schema to be created"
}

variable "db_name" {
  type        = string
  description = "The name of the database"
}

variable "admin_role" {
  type        = string
  description = "The name of the admin PostgreSQL role for the DB"
}

variable "app_role" {
  type        = string
  description = "The name of the app PostgreSQL role for the DB"
}
