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
	"fmt"

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
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Importer: &schema.ResourceImporter{
			// NOTE: A pure pass through is insufficient if `books_count` gets added.
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceAuthorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, diags := getClientFromMeta(meta)
	if diags != nil {
		return diags
	}

	firstName, ok := d.Get("first_name").(string)
	if !ok {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not determine author first name",
			Detail:   "Invalid first name parameter type",
		})
		return diags
	}
	lastName, ok := d.Get("last_name").(string)
	if !ok {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not determine author last name",
			Detail:   "Invalid last name parameter type",
		})
		return diags
	}

	a := booksclient.Author{FirstName: firstName, LastName: lastName}
	aar, err := c.AddAuthor(ctx, a)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%d", aar.AuthorID))
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

	gar := booksclient.GetAuthorRequest{AuthorID: id}
	a, err := c.GetAuthor(ctx, gar)
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

	return nil
}

func resourceAuthorUpdate(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	diags := []diag.Diagnostic{
		{
			Severity: diag.Error,
			Summary:  "Author cannot be changed after creation",
			Detail:   "Unsupported operation",
		},
	}
	return diags
}

func resourceAuthorDelete(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	diags := []diag.Diagnostic{
		{
			Severity: diag.Error,
			Summary:  "Author cannot be changed after creation",
			Detail:   "Unsupported operation",
		},
	}
	return diags
}
