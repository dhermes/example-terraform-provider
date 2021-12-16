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
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/dhermes/example-terraform-provider/pkg/booksclient"
)

const (
	dateLayout = "2006-01-02"
)

// resourceBook returns the `book` resource in the Terraform provider for
// the Books API.
func resourceBook() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBookCreate,
		ReadContext:   resourceBookRead,
		UpdateContext: resourceBookUpdate,
		DeleteContext: resourceBookDelete,
		Schema: map[string]*schema.Schema{
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"author_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"publish_date": {
				Type:     schema.TypeString,
				Required: true,
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

func resourceBookCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, diags := getClientFromMeta(meta)
	if diags != nil {
		return diags
	}

	authorIDStr, ok := d.Get("author_id").(string)
	if !ok {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not determine author ID",
			Detail:   "Invalid author ID parameter type",
		})
		return diags
	}
	title, ok := d.Get("title").(string)
	if !ok {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not determine book title",
			Detail:   "Invalid book title parameter type",
		})
		return diags
	}
	publishDateStr, ok := d.Get("publish_date").(string)
	if !ok {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not determine book publish date",
			Detail:   "Invalid book publish date parameter type",
		})
		return diags
	}

	authorID, err := uuid.Parse(authorIDStr)
	if err != nil {
		return diag.FromErr(err)
	}
	publishDate, err := time.Parse(dateLayout, publishDateStr)
	if err != nil {
		return diag.FromErr(err)
	}
	publishDate = publishDate.UTC()

	b := booksclient.Book{Title: title, AuthorID: authorID, PublishDate: &publishDate}
	abr, err := c.AddBook(ctx, b)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(abr.BookID.String())
	return resourceBookRead(ctx, d, meta)
}

func resourceBookRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, diags := getClientFromMeta(meta)
	if diags != nil {
		return diags
	}

	idStr := d.Id()
	id, diags := idFromString(idStr)
	if diags != nil {
		return diags
	}

	gbr := booksclient.GetBookRequest{BookID: id}
	b, err := c.GetBook(ctx, gbr)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("title", b.Title)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("author_id", b.AuthorID.String())
	if err != nil {
		return diag.FromErr(err)
	}

	if b.PublishDate != nil {
		err = d.Set("publish_date", b.PublishDate.Format(dateLayout))
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceBookUpdate(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	diags := []diag.Diagnostic{
		{
			Severity: diag.Error,
			Summary:  "Book cannot be changed after creation",
			Detail:   "Unsupported operation",
		},
	}
	return diags
}

func resourceBookDelete(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	diags := []diag.Diagnostic{
		{
			Severity: diag.Error,
			Summary:  "Book cannot be changed after creation",
			Detail:   "Unsupported operation",
		},
	}
	return diags
}
