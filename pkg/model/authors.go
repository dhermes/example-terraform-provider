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

package model

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const (
	insertAuthor = `
INSERT INTO
  authors (id, first_name, last_name)
VALUES
  ($1, $2, $3)
`
	getAuthorByID = `
SELECT
  id,
  first_name,
  last_name,
  (
    SELECT
      COUNT(*) AS book_count
    FROM
      books
    WHERE
      author_id = $1
  )
FROM
  authors
WHERE
  id = $1
`
	getAllAuthors = `
SELECT
  a.id, a.first_name, a.last_name, COALESCE(b.book_count, 0)
FROM
  authors AS a
FULL OUTER JOIN (
  SELECT
    author_id, COUNT(*) AS book_count
  FROM
    books
  GROUP BY
    author_id
) AS b
ON
  a.id = b.author_id
`
)

// InsertAuthor inserts an author into the database.
func InsertAuthor(ctx context.Context, pool *sql.DB, a Author) (uuid.UUID, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return uuid.Nil, err
	}

	_, err = pool.ExecContext(ctx, insertAuthor, id, a.FirstName, a.LastName)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

// GetAuthorByID gets an authors from the database by ID.
func GetAuthorByID(ctx context.Context, pool *sql.DB, id uuid.UUID) (*Author, error) {
	row := pool.QueryRowContext(ctx, getAuthorByID, id)

	a := Author{}
	err := row.Scan(&a.ID, &a.FirstName, &a.LastName, &a.BookCount)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

// GetAllAuthors gets (all) authors from the database.
//
// This function does not use paging (but it would in a real application).
func GetAllAuthors(ctx context.Context, pool *sql.DB) ([]Author, error) {
	rows, err := pool.QueryContext(ctx, getAllAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	authors := []Author{}
	for rows.Next() {
		a := Author{}
		err = rows.Scan(&a.ID, &a.FirstName, &a.LastName, &a.BookCount)
		if err != nil {
			return nil, err
		}
		authors = append(authors, a)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return authors, nil
}
