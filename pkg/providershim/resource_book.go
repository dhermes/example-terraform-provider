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

// resourceBook returns the `books_api_book` resource in the Terraform provider
// for the Books API.
func resourceBook() *schema.Resource {
	var stub *booksprovider.ResourceBook
	return &schema.Resource{
		CreateContext: resourceBookCreate,
		ReadContext:   resourceBookRead,
		UpdateContext: resourceBookUpdate,
		DeleteContext: resourceBookDelete,
		Schema:        stub.Schema(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceBookCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, diags := getClientFromMeta(meta)
	if diags != nil {
		return diags
	}

	rb, err := booksprovider.NewResourceBook(d)
	if err != nil {
		return terraform.AppendDiagnostic(err, nil)
	}

	err = rb.Create(ctx, c)
	return terraform.AppendDiagnostic(err, nil)
}

func resourceBookRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, diags := getClientFromMeta(meta)
	if diags != nil {
		return diags
	}

	rb, err := booksprovider.NewResourceBook(d)
	if err != nil {
		return terraform.AppendDiagnostic(err, nil)
	}

	err = rb.Read(ctx, c)
	return terraform.AppendDiagnostic(err, nil)
}

func resourceBookUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, diags := getClientFromMeta(meta)
	if diags != nil {
		return diags
	}

	rb, err := booksprovider.NewResourceBook(d)
	if err != nil {
		return terraform.AppendDiagnostic(err, nil)
	}

	err = rb.Update(ctx, c)
	return terraform.AppendDiagnostic(err, nil)
}

func resourceBookDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, diags := getClientFromMeta(meta)
	if diags != nil {
		return diags
	}

	rb, err := booksprovider.NewResourceBook(d)
	if err != nil {
		return terraform.AppendDiagnostic(err, nil)
	}

	err = rb.Delete(ctx, c)
	return terraform.AppendDiagnostic(err, nil)
}
