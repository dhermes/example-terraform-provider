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

data "books_api_author" "vonnegut" {
  provider = books.local

  first_name = "Kurt"
  last_name  = "Vonnegut"
}

resource "books_api_book" "sirens" {
  provider = books.local

  author_id    = data.books_api_author.vonnegut.id
  title        = "The Sirens of Titan"
  publish_date = "1959-01-01"
}
