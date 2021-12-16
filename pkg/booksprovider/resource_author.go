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

// ResourceAuthor represents a `books_api_author` resource and directly
// interacts with the Terraform plugin SDK to check for new or updated
// values.
type ResourceAuthor struct {
	mutex sync.RWMutex
	d     *schema.ResourceData

	FirstName *string    `terraform:"first_name,required"`
	LastName  *string    `terraform:"last_name,required"`
	BookCount *int       `terraform:"book_count,computed"`
	ID        *uuid.UUID `terraform:"id,string,computed"`
}

// Create is the create (C) component of the CRUD lifecycle for
// the `books_api_author` resource.
func (ra *ResourceAuthor) Create(ctx context.Context, c booksclient.Client) error {
	err := ra.Populate()
	if err != nil {
		return err
	}

	a := booksclient.Author{FirstName: ra.GetFirstName(), LastName: ra.GetLastName()}
	aar, err := c.AddAuthor(ctx, a)
	if err != nil {
		return err
	}

	ra.ID = &aar.AuthorID
	err = ra.Persist()
	if err != nil {
		return err
	}

	return ra.Read(ctx, c)
}

// Read is the read (R) component of the CRUD lifecycle for
// the `books_api_author` resource.
func (ra *ResourceAuthor) Read(ctx context.Context, c booksclient.Client) error {
	err := ra.Populate()
	if err != nil {
		return err
	}

	id := ra.GetID()
	gabir := booksclient.GetAuthorByIDRequest{AuthorID: id}
	a, err := c.GetAuthorByID(ctx, gabir)
	if err != nil {
		return err
	}

	ra.FirstName = &a.FirstName
	ra.LastName = &a.LastName
	bc := int(a.BookCount)
	ra.BookCount = &bc
	return ra.Persist()
}

// Update is the update (U) component of the CRUD lifecycle for
// the `books_api_author` resource.
func (ra *ResourceAuthor) Update(ctx context.Context, c booksclient.Client) error {
	if !ra.Changed() {
		return ra.Read(ctx, c)
	}

	err := ra.Populate()
	if err != nil {
		return err
	}

	a := booksclient.Author{ID: ra.ID, FirstName: ra.GetFirstName(), LastName: ra.GetLastName()}
	_, err = c.UpdateAuthor(ctx, a)
	if err != nil {
		return err
	}

	return ra.Read(ctx, c)
}

// Delete is the delete (D) component of the CRUD lifecycle for
// the `books_api_author` resource.
func (ra *ResourceAuthor) Delete(ctx context.Context, c booksclient.Client) error {
	err := ra.Populate()
	if err != nil {
		return err
	}

	id := ra.GetID()
	dar := booksclient.DeleteAuthorRequest{AuthorID: id}
	_, err = c.DeleteAuthorByID(ctx, dar)
	if err != nil {
		return err
	}

	return nil
}
