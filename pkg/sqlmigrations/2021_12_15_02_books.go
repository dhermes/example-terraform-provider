// Copyright 2021 Danny Hermes
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sqlmigrations

import (
	"context"
	"database/sql"

	"github.com/dhermes/golembic"
)

// NOTE: Ensure that
//       * `AddBooksTable` satisfies `golembic.UpMigration`.
var (
	_ golembic.UpMigration = AddBooksTable
)

const (
	booksCreate = `
CREATE TABLE books (
  id UUID NOT NULL,
  author_id UUID NOT NULL,
  title TEXT NOT NULL,
  publish_date DATE NOT NULL
)
`
	booksPK = `
ALTER TABLE
  books
ADD CONSTRAINT
  pk_books_id
PRIMARY KEY
  (id)
`
)

// AddBooksTable runs SQL statements required for adding the `books` table.
func AddBooksTable(ctx context.Context, tx *sql.Tx) error {
	err := applySQL(ctx, tx, booksCreate)
	if err != nil {
		return err
	}

	return applySQL(ctx, tx, booksPK)
}
