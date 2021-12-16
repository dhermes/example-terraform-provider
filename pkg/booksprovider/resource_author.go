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

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/dhermes/example-terraform-provider/pkg/booksclient"
	"github.com/dhermes/example-terraform-provider/pkg/terraform"
)

// ResourceAuthor represents a `books_api_author` resource and directly
// interacts with the Terraform plugin SDK to check for new or updated
// values.
type ResourceAuthor struct {
	d *schema.ResourceData

	FirstName *string    `terraform:"first_name,required"`
	LastName  *string    `terraform:"last_name,required"`
	BookCount *int       `terraform:"book_count,computed"`
	ID        *uuid.UUID `terraform:"id,string,computed"`
}

// NewResourceAuthor creates a new `ResourceAuthor` from a Terraform resource
// data struct.
func NewResourceAuthor(d *schema.ResourceData) ResourceAuthor {
	return ResourceAuthor{d: d}
}

// GetFirstName is a value accessor for a pointer field; a safe dereference.
// (The goal is to make code that can be autogenerated.)
func (ra *ResourceAuthor) GetFirstName() string {
	p := ra.FirstName
	if p == nil {
		return ""
	}
	return *p
}

// GetLastName is a value accessor for a pointer field; a safe dereference.
// (The goal is to make code that can be autogenerated.)
func (ra *ResourceAuthor) GetLastName() string {
	p := ra.LastName
	if p == nil {
		return ""
	}
	return *p
}

// Populate populates the fields in this struct based on the `terraform`
// struct tags. (The goal is to make code that can be autogenerated.)
func (ra *ResourceAuthor) Populate() error {
	var err error

	firstName, ok := ra.d.Get("first_name").(string)
	if !ok {
		err = terraform.DiagnosticError{
			Summary: "Could not determine author first name",
			Detail:  "Invalid first name parameter type",
		}
		return err
	}

	lastName, ok := ra.d.Get("last_name").(string)
	if !ok {
		err = terraform.DiagnosticError{
			Summary: "Could not determine author last name",
			Detail:  "Invalid last name parameter type",
		}
		return err
	}

	idStr := ra.d.Id()
	id := ra.ID
	if idStr != "" {
		parsed, err := idFromString(idStr)
		if err != nil {
			return err
		}
		id = &parsed
	}

	ra.FirstName = &firstName
	ra.LastName = &lastName
	ra.ID = id
	return nil
}

// Persist writes back fields in this struct to the Terraform resource data
// struct based on the `terraform` struct tags. (The goal is to make code that
// can be autogenerated.)
func (ra *ResourceAuthor) Persist() error {
	ra.d.SetId(ra.ID.String())
	return nil
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
