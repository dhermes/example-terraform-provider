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
	"context"
)

// Client is an abstraction for the Books API.
//
// The books service helps to build a catalog of books along with managing the
// authors of those books. Once the catalog is populated, the authors can be
// listed and books can be queried by author.
type Client interface {
	AddAuthor(context.Context, Author) (*AddAuthorResponse, error)
	UpdateAuthor(context.Context, Author) (*Empty, error)
	GetAuthorByID(context.Context, GetAuthorByIDRequest) (*Author, error)
	GetAuthorByName(context.Context, GetAuthorByNameRequest) (*Author, error)
	GetAuthors(context.Context, Empty) (*GetAuthorsResponse, error)
	DeleteAuthorByID(context.Context, DeleteAuthorRequest) (*Empty, error)

	AddBook(context.Context, Book) (*AddBookResponse, error)
	UpdateBook(context.Context, Book) (*Empty, error)
	GetBookByID(context.Context, GetBookByIDRequest) (*Book, error)
	GetBooks(context.Context, GetBooksRequest) (*GetBooksResponse, error)
	DeleteBookByID(context.Context, DeleteBookRequest) (*Empty, error)
}
