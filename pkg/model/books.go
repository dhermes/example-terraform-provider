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
	insertBook = `
INSERT INTO
  books (id, author_id, title, publish_date)
SELECT
  $1, $2, $3, $4
WHERE
  EXISTS (
    SELECT 1 FROM authors AS a WHERE a.id = $2 FOR UPDATE
  )
`
	updateBook = `
UPDATE
  books
SET
  author_id = $2,
  title = $3,
  publish_date = $4
WHERE
  id = $1
`
	getBookByID = `
SELECT
  id, author_id, title, publish_date
FROM
  books
WHERE
  id = $1
`
	getAllBooksByAuthor = `
SELECT
  id, author_id, title, publish_date
FROM
  books
WHERE
  author_id = $1
`
	deleteBookByID = `
DELETE FROM
  books
WHERE
  id = $1
`
)

// InsertBook inserts a book into the database.
func InsertBook(ctx context.Context, pool *sql.DB, b Book) (uuid.UUID, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return uuid.Nil, err
	}

	// NOTE: Instead of doing two round trips, the `insertBook` makes sure
	//       the author ID exists via a subquery. This is effectively the
	//       same cost as using a foreign key, but does not **require** the
	//       use of a foreign key.
	_, err = pool.ExecContext(ctx, insertBook, id, b.AuthorID, b.Title, b.PublishDate)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

// UpdateBook updates a book from the database directly by ID.
func UpdateBook(ctx context.Context, pool *sql.DB, b Book) error {
	result, err := pool.ExecContext(ctx, updateBook, b.ID, b.AuthorID, b.Title, b.PublishDate)
	if err != nil {
		return err
	}

	updateCount, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if updateCount == 0 {
		return errors.New("could not update book, does not exist")
	}

	return nil
}

// GetBookByID gets a book from the database by ID.
func GetBookByID(ctx context.Context, pool *sql.DB, id uuid.UUID) (*Book, error) {
	row := pool.QueryRowContext(ctx, getBookByID, id)

	b := Book{}
	err := row.Scan(&b.ID, &b.AuthorID, &b.Title, &b.PublishDate)
	if err != nil {
		return nil, err
	}

	return &b, nil
}

// GetAllBooksByAuthor gets (all) books by an author from the database.
//
// This function does not use paging (but it would in a real application).
func GetAllBooksByAuthor(ctx context.Context, pool *sql.DB, authorID uuid.UUID) ([]Book, error) {
	rows, err := pool.QueryContext(ctx, getAllBooksByAuthor, authorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []Book{}
	for rows.Next() {
		b := Book{}
		err = rows.Scan(&b.ID, &b.AuthorID, &b.Title, &b.PublishDate)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return books, nil
}

// DeleteBookByID deletes a book from the database by ID.
func DeleteBookByID(ctx context.Context, pool *sql.DB, id uuid.UUID) error {
	result, err := pool.ExecContext(ctx, deleteBookByID, id)
	if err != nil {
		return err
	}

	deleteCount, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if deleteCount == 0 {
		return errors.New("could not delete book, does not exist")
	}

	return nil
}
