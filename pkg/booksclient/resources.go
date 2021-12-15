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
	"time"
)

// Empty is an empty struct used for RPC inputs / outputs.
type Empty struct{}

// Author contains information about an author.
//
// At creation time, neither the ID nor the book count can be set.
type Author struct {
	// FirstName is the given name of the author.
	FirstName string `json:"first_name"`
	// LastName is the surname of the author.
	LastName string `json:"last_name"`
	// ID is the database identifier, if the author has already been created.
	ID uint64 `json:"id,string,omitempty"`
	// BookCount is the number of books by the author in the books service.
	BookCount uint32 `json:"book_count,omitempty"`
}

// Book contains information about a book.
//
// At creation time, the ID cannot be set.
type Book struct {
	// Title is the book title.
	Title string `json:"title"`
	// AuthorID is the ID of the author of the book.
	AuthorID uint64 `json:"author_id,string"`
	// PublishDate is the date the book was published.
	PublishDate *time.Time `json:"publish_date,omitempty"`
	// ID is the database identifier, if the book has already been created.
	ID uint64 `json:"id,string,omitempty"`
}
