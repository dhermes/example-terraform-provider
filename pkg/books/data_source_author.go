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

// dataSourceAuthor returns the `author` data source in the Terraform provider
// for the Books API.
func dataSourceAuthor() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAuthorRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"first_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_name": {
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

	idStr, ok := d.Get("id").(string)
	if !ok {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not determine author ID",
			Detail:   "Invalid ID parameter type",
		})
		return diags
	}
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

	d.SetId(idStr)
	return diags
}
