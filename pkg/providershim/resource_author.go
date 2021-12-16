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

	"github.com/dhermes/example-terraform-provider/pkg/booksprovider"
	"github.com/dhermes/example-terraform-provider/pkg/terraform"
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

	err := booksprovider.ResourceAuthorCreate(ctx, d, c)
	return terraform.AppendDiagnostic(err, nil)
}

func resourceAuthorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, diags := getClientFromMeta(meta)
	if diags != nil {
		return diags
	}

	err := booksprovider.ResourceAuthorRead(ctx, d, c)
	return terraform.AppendDiagnostic(err, nil)
}

func resourceAuthorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, diags := getClientFromMeta(meta)
	if diags != nil {
		return diags
	}

	err := booksprovider.ResourceAuthorUpdate(ctx, d, c)
	return terraform.AppendDiagnostic(err, nil)
}

func resourceAuthorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, diags := getClientFromMeta(meta)
	if diags != nil {
		return diags
	}

	err := booksprovider.ResourceAuthorDelete(ctx, d, c)
	return terraform.AppendDiagnostic(err, nil)
}
