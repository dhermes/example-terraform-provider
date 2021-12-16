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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/dhermes/example-terraform-provider/pkg/booksclient"
)

// ResourceAuthor represents a `books_api_author` resource and directly
// interacts with the Terraform plugin SDK to check for new or updated
// values.
type ResourceAuthor struct {
	d *schema.ResourceData
}

// NewResourceAuthor creates a new `ResourceAuthor` from a Terraform resource
// data struct.
func NewResourceAuthor(d *schema.ResourceData) ResourceAuthor {
	return ResourceAuthor{d: d}
}

// Create is the create (C) component of the CRUD lifecycle for
// the `books_api_author` resource.
func (ra *ResourceAuthor) Create(ctx context.Context, c booksclient.Client) error {
	a, err := authorFromResourceData(ra.d)
	if err != nil {
		return err
	}

	aar, err := c.AddAuthor(ctx, *a)
	if err != nil {
		return err
	}

	ra.d.SetId(aar.AuthorID.String())
	return ra.Read(ctx, c)
}

// Read is the read (R) component of the CRUD lifecycle for
// the `books_api_author` resource.
func (ra *ResourceAuthor) Read(ctx context.Context, c booksclient.Client) error {
	idStr := ra.d.Id()
	id, err := idFromString(idStr)
	if err != nil {
		return err
	}

	gabir := booksclient.GetAuthorByIDRequest{AuthorID: id}
	a, err := c.GetAuthorByID(ctx, gabir)
	if err != nil {
		return err
	}

	err = ra.d.Set("first_name", a.FirstName)
	if err != nil {
		return err
	}

	err = ra.d.Set("last_name", a.LastName)
	if err != nil {
		return err
	}

	err = ra.d.Set("book_count", int(a.BookCount))
	if err != nil {
		return err
	}

	return nil
}

// Update is the update (U) component of the CRUD lifecycle for
// the `books_api_author` resource.
func (ra *ResourceAuthor) Update(ctx context.Context, c booksclient.Client) error {
	anyChange := ra.d.HasChange("first_name") || ra.d.HasChange("last_name")
	if !anyChange {
		return ra.Read(ctx, c)
	}

	a, err := authorFromResourceData(ra.d)
	if err != nil {
		return err
	}

	_, err = c.UpdateAuthor(ctx, *a)
	if err != nil {
		return err
	}

	return ra.Read(ctx, c)
}

// Delete is the delete (D) component of the CRUD lifecycle for
// the `books_api_author` resource.
func (ra *ResourceAuthor) Delete(ctx context.Context, c booksclient.Client) error {
	idStr := ra.d.Id()
	id, err := idFromString(idStr)
	if err != nil {
		return err
	}

	dar := booksclient.DeleteAuthorRequest{AuthorID: id}
	_, err = c.DeleteAuthorByID(ctx, dar)
	if err != nil {
		return err
	}

	// This is superfluous but added here for explicitness.
	ra.d.SetId("")
	return nil
}
