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

package booksprovider

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/dhermes/example-terraform-provider/pkg/booksclient"
)

// ResourceBook represents a `books_api_book` resource and directly
// interacts with the Terraform plugin SDK to check for new or updated
// values.
type ResourceBook struct {
	mutex sync.RWMutex
	d     *schema.ResourceData

	Title       *string    `terraform:"title,required"`
	AuthorID    *uuid.UUID `terraform:"author_id,string,required"`
	PublishDate *Date      `terraform:"publish_date,string,required"`
	ID          *uuid.UUID `terraform:"id,string,computed"`
}

// Create is the create (C) component of the CRUD lifecycle for
// the `books_api_book` resource.
//
// NOTE: This assumes the called has already invoked `rb.Populate()`, either
//       directly or indirectly, e.g. via `NewResourceBook()`.
func (rb *ResourceBook) Create(ctx context.Context, c booksclient.Client) error {
	pd := rb.PublishDate
	b := booksclient.Book{Title: rb.GetTitle(), AuthorID: rb.GetAuthorID(), PublishDate: &pd.Time}
	abr, err := c.AddBook(ctx, b)
	if err != nil {
		return err
	}

	rb.ID = &abr.BookID
	err = rb.Persist()
	if err != nil {
		return err
	}

	return rb.Read(ctx, c)
}

// Read is the read (R) component of the CRUD lifecycle for
// the `books_api_book` resource.
//
// NOTE: This assumes the called has already invoked `rb.Populate()`, either
//       directly or indirectly, e.g. via `NewResourceBook()`.
func (rb *ResourceBook) Read(ctx context.Context, c booksclient.Client) error {
	id := rb.GetID()
	gbbir := booksclient.GetBookByIDRequest{BookID: id}
	b, err := c.GetBookByID(ctx, gbbir)
	if err != nil {
		return err
	}

	rb.Title = &b.Title
	rb.AuthorID = &b.AuthorID
	if b.PublishDate != nil {
		rb.PublishDate = &Date{Time: *b.PublishDate}
	}
	return rb.Persist()
}

// Update is the update (U) component of the CRUD lifecycle for
// the `books_api_book` resource.
//
// NOTE: This assumes the called has already invoked `rb.Populate()`, either
//       directly or indirectly, e.g. via `NewResourceBook()`.
func (rb *ResourceBook) Update(ctx context.Context, c booksclient.Client) error {
	if !rb.Changed() {
		return rb.Read(ctx, c)
	}

	pd := rb.PublishDate
	b := booksclient.Book{ID: rb.ID, Title: rb.GetTitle(), AuthorID: rb.GetAuthorID(), PublishDate: &pd.Time}
	_, err := c.UpdateBook(ctx, b)
	if err != nil {
		return err
	}

	return rb.Read(ctx, c)
}

// Delete is the delete (D) component of the CRUD lifecycle for
// the `books_api_book` resource.
//
// NOTE: This assumes the called has already invoked `rb.Populate()`, either
//       directly or indirectly, e.g. via `NewResourceBook()`.
func (rb *ResourceBook) Delete(ctx context.Context, c booksclient.Client) error {
	id := rb.GetID()
	dbr := booksclient.DeleteBookRequest{BookID: id}
	_, err := c.DeleteBookByID(ctx, dbr)
	return err
}
