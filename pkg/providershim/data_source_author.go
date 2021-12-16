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

package providershim

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/dhermes/example-terraform-provider/pkg/booksclient"
)

// dataSourceAuthor returns the `author` data source in the Terraform provider
// for the Books API.
func dataSourceAuthor() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAuthorRead,
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
	}
}

func dataSourceAuthorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	gabnr := booksclient.GetAuthorByNameRequest{FirstName: firstName, LastName: lastName}
	a, err := c.GetAuthorByName(ctx, gabnr)
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

	d.SetId(a.ID.String())
	return diags
}