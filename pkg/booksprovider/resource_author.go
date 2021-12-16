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

// ResourceAuthorCreate is the create (C) component of the CRUD lifecycle for
// the `books_api_author` resource.
func ResourceAuthorCreate(ctx context.Context, d *schema.ResourceData, c booksclient.Client) error {
	a, err := authorFromResourceData(d)
	if err != nil {
		return err
	}

	aar, err := c.AddAuthor(ctx, *a)
	if err != nil {
		return err
	}

	d.SetId(aar.AuthorID.String())
	return ResourceAuthorRead(ctx, d, c)
}

// ResourceAuthorRead is the read (R) component of the CRUD lifecycle for
// the `books_api_author` resource.
func ResourceAuthorRead(ctx context.Context, d *schema.ResourceData, c booksclient.Client) error {
	idStr := d.Id()
	id, err := idFromString(idStr)
	if err != nil {
		return err
	}

	gabir := booksclient.GetAuthorByIDRequest{AuthorID: id}
	a, err := c.GetAuthorByID(ctx, gabir)
	if err != nil {
		return err
	}

	err = d.Set("first_name", a.FirstName)
	if err != nil {
		return err
	}

	err = d.Set("last_name", a.LastName)
	if err != nil {
		return err
	}

	err = d.Set("book_count", int(a.BookCount))
	if err != nil {
		return err
	}

	return nil
}

// ResourceAuthorUpdate is the update (U) component of the CRUD lifecycle for
// the `books_api_author` resource.
func ResourceAuthorUpdate(ctx context.Context, d *schema.ResourceData, c booksclient.Client) error {
	anyChange := d.HasChange("first_name") || d.HasChange("last_name")
	if !anyChange {
		return ResourceAuthorRead(ctx, d, c)
	}

	a, err := authorFromResourceData(d)
	if err != nil {
		return err
	}

	_, err = c.UpdateAuthor(ctx, *a)
	if err != nil {
		return err
	}

	return ResourceAuthorRead(ctx, d, c)
}

// ResourceAuthorDelete is the delete (D) component of the CRUD lifecycle for
// the `books_api_author` resource.
func ResourceAuthorDelete(ctx context.Context, d *schema.ResourceData, c booksclient.Client) error {
	idStr := d.Id()
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
	d.SetId("")
	return nil
}
