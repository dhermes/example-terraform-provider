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
//       * `AddAuthorsTable` satisfies `golembic.UpMigration`.
var (
	_ golembic.UpMigration = AddAuthorsTable
)

const (
	authorsCreate = `
CREATE TABLE authors (
  id UUID NOT NULL,
  first_name TEXT NOT NULL,
  last_name TEXT NOT NULL
)
`
	authorsPK = `
ALTER TABLE
  authors
ADD CONSTRAINT
  pk_authors_id
PRIMARY KEY
  (id)
`
)

// AddAuthorsTable runs SQL statements required for adding the `authors` table.
func AddAuthorsTable(ctx context.Context, tx *sql.Tx) error {
	err := applySQL(ctx, tx, authorsCreate)
	if err != nil {
		return err
	}

	return applySQL(ctx, tx, authorsPK)
}
