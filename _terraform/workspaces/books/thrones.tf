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

resource "books_api_author" "grr_martin" {
  provider = books.local

  first_name = "George R.R."
  last_name  = "Martin"
}

resource "books_api_book" "song_fire_ice1" {
  provider = books.local

  author_id    = books_api_author.grr_martin.id
  title        = "A Game of Thrones"
  publish_date = "1996-08-01"
}

resource "books_api_book" "song_fire_ice2" {
  provider = books.local

  author_id    = books_api_author.grr_martin.id
  title        = "A Clash of Kings"
  publish_date = "1998-11-16"
}

resource "books_api_book" "song_fire_ice3" {
  provider = books.local

  author_id    = books_api_author.grr_martin.id
  title        = "A Storm of Swords"
  publish_date = "2000-08-08"
}

resource "books_api_book" "song_fire_ice4" {
  provider = books.local

  author_id    = books_api_author.grr_martin.id
  title        = "A Feast for Crows"
  publish_date = "2005-10-17"
}

resource "books_api_book" "song_fire_ice5" {
  provider = books.local

  author_id    = books_api_author.grr_martin.id
  title        = "A Dance with Dragons"
  publish_date = "2011-07-12"
}
