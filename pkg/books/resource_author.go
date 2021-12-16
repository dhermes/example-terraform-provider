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

package books

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/dhermes/example-terraform-provider/pkg/booksclient"
)

// resourceAuthor returns the `author` resource in the Terraform provider for
// the Books API.
func resourceAuthor() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAuthorCreate,
		ReadContext:   resourceAuthorRead,
		UpdateContext: resourceAuthorUpdate,
		DeleteContext: resourceAuthorDelete,
		Schema: map[string]*schema.Schema{
			"first_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"book_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceAuthorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, diags := getClientFromMeta(meta)
	if diags != nil {
		return diags
	}

	a, diags := authorFromResourceData(d)
	if diags != nil {
		return diags
	}

	aar, err := c.AddAuthor(ctx, *a)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(aar.AuthorID.String())
	return resourceAuthorRead(ctx, d, meta)
}

func resourceAuthorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, diags := getClientFromMeta(meta)
	if diags != nil {
		return diags
	}

	idStr := d.Id()
	id, diags := idFromString(idStr)
	if diags != nil {
		return diags
	}

	gabir := booksclient.GetAuthorByIDRequest{AuthorID: id}
	a, err := c.GetAuthorByID(ctx, gabir)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("first_name", a.FirstName)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("last_name", a.LastName)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("book_count", int(a.BookCount))
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceAuthorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	anyChange := d.HasChange("first_name") || d.HasChange("last_name")
	if !anyChange {
		return resourceAuthorRead(ctx, d, meta)
	}

	c, diags := getClientFromMeta(meta)
	if diags != nil {
		return diags
	}

	a, diags := authorFromResourceData(d)
	if diags != nil {
		return diags
	}

	_, err := c.UpdateAuthor(ctx, *a)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceAuthorRead(ctx, d, meta)
}

func resourceAuthorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, diags := getClientFromMeta(meta)
	if diags != nil {
		return diags
	}

	idStr := d.Id()
	id, diags := idFromString(idStr)
	if diags != nil {
		return diags
	}

	dar := booksclient.DeleteAuthorRequest{AuthorID: id}
	_, err := c.DeleteAuthorByID(ctx, dar)
	if err != nil {
		return diag.FromErr(err)
	}

	// This is superfluous but added here for explicitness.
	d.SetId("")
	return nil
}

func authorFromResourceData(d *schema.ResourceData) (*booksclient.Author, diag.Diagnostics) {
	var diags diag.Diagnostics

	firstName, ok := d.Get("first_name").(string)
	if !ok {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not determine author first name",
			Detail:   "Invalid first name parameter type",
		})
		return nil, diags
	}
	lastName, ok := d.Get("last_name").(string)
	if !ok {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not determine author last name",
			Detail:   "Invalid last name parameter type",
		})
		return nil, diags
	}

	a := booksclient.Author{FirstName: firstName, LastName: lastName}
	idStr := d.Id()
	if idStr != "" {
		id, diags := idFromString(idStr)
		if diags != nil {
			return nil, diags
		}
		a.ID = &id
	}

	return &a, nil
}
