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

package booksclient

import (
	"github.com/google/uuid"
)

// AddAuthorResponse is the response after an author was added.
type AddAuthorResponse struct {
	// AuthorID is the ID of the newly created author.
	AuthorID uuid.UUID `json:"author_id"`
}

// GetAuthorByIDRequest is the request for a query to get an author by ID.
type GetAuthorByIDRequest struct {
	// AuthorID is the ID of the author being queried.
	AuthorID uuid.UUID `json:"author_id"`
}

// GetAuthorByNameRequest is the request for a query to get an author by name.
type GetAuthorByNameRequest struct {
	// FirstName is the given name of the author.
	FirstName string `json:"first_name"`
	// LastName is the surname of the author.
	LastName string `json:"last_name"`
}

// GetAuthorsResponse is the response for a list authors query.
type GetAuthorsResponse struct {
	// Authors is the sequence of retrieved authors.
	Authors []Author `json:"authors"`
}

// DeleteAuthorRequest is the request for an author deletion.
type DeleteAuthorRequest struct {
	// AuthorID is the ID of the author being deleted.
	AuthorID uuid.UUID `json:"author_id"`
}

// AddBookResponse is the response after a book was added.
type AddBookResponse struct {
	// BookID is the ID of the newly created book.
	BookID uuid.UUID `json:"book_id"`
}

// GetBookByIDRequest is the request for a query to get a book by ID.
type GetBookByIDRequest struct {
	// BookID is the ID of the book being queried.
	BookID uuid.UUID `json:"book_id"`
}

// GetBooksRequest is a request for a books-by-author query.
type GetBooksRequest struct {
	// AuthorID is the ID of the author of the books.
	AuthorID uuid.UUID `json:"author_id"`
}

// GetBooksResponse is the response for a books query.
type GetBooksResponse struct {
	// Books is the sequence of retrieved books.
	Books []Book `json:"books"`
}

// DeleteBookRequest is the request for a book deletion.
type DeleteBookRequest struct {
	// BookID is the ID of the book being deleted.
	BookID uuid.UUID `json:"book_id"`
}
