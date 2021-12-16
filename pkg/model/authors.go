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
	"errors"

	"github.com/google/uuid"
)

const (
	insertAuthor = `
INSERT INTO
  authors (id, first_name, last_name)
VALUES
  ($1, $2, $3)
`
	updateAuthor = `
UPDATE
  authors
SET
  first_name = $2,
  last_name = $3
WHERE
  id = $1
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
      books AS b
    WHERE
      b.author_id = a.id
  )
FROM
  authors AS a
WHERE
  id = $1
`
	getAuthorByName = `
SELECT
  id,
  first_name,
  last_name,
  (
    SELECT
      COUNT(*) AS book_count
    FROM
      books AS b
    WHERE
      b.author_id = a.id
  )
FROM
  authors AS a
WHERE
  first_name = $1 AND
  last_name = $2
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
	deleteAuthorByID = `
DELETE FROM
  authors AS a
WHERE
  a.id = $1 AND
  NOT EXISTS (
    SELECT 1 FROM books AS b WHERE b.author_id = a.id FOR UPDATE
  )
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

// UpdateAuthor updates an author from the database directly by ID.
func UpdateAuthor(ctx context.Context, pool *sql.DB, a Author) error {
	result, err := pool.ExecContext(ctx, updateAuthor, a.ID, a.FirstName, a.LastName)
	if err != nil {
		return err
	}

	updateCount, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if updateCount == 0 {
		return errors.New("could not update author, does not exist")
	}

	return nil
}

// GetAuthorByID gets an author from the database by ID.
func GetAuthorByID(ctx context.Context, pool *sql.DB, id uuid.UUID) (*Author, error) {
	row := pool.QueryRowContext(ctx, getAuthorByID, id)

	a := Author{}
	err := row.Scan(&a.ID, &a.FirstName, &a.LastName, &a.BookCount)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

// GetAuthorByName gets an author from the database by name.
func GetAuthorByName(ctx context.Context, pool *sql.DB, firstName, lastName string) (*Author, error) {
	row := pool.QueryRowContext(ctx, getAuthorByName, firstName, lastName)

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

// DeleteAuthorByID deletes an author from the database by ID.
func DeleteAuthorByID(ctx context.Context, pool *sql.DB, id uuid.UUID) error {
	result, err := pool.ExecContext(ctx, deleteAuthorByID, id)
	if err != nil {
		return err
	}

	deleteCount, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if deleteCount == 0 {
		return errors.New("could not delete author, does not exist or still has books")
	}

	return nil
}
