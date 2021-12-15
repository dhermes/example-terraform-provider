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
	"net/url"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/dhermes/example-terraform-provider/pkg/booksclient"
)

// Provider returns a Terraform provider for the Books API.
//
// It composes all of the resources and data sources as well as
// provider arguments into one struct.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"addr": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("BOOKS_API_ADDR", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"books_api_author": resourceAuthor(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"books_api_author": dataSourceAuthor(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(_ context.Context, d *schema.ResourceData) (meta interface{}, diags diag.Diagnostics) {
	addr, ok := d.Get("addr").(string)
	if addr == "" || !ok {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Books API client",
			Detail:   "Unable to determine address for Books API",
		})
		return
	}

	_, err := url.Parse(addr)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Books API client",
			Detail:   "Failed to parse Books API base address as a URL",
		})
		return
	}

	c, err := booksclient.NewHTTPClient(booksclient.OptAddr(addr))
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Books API client",
			Detail:   "Unable to create Books API client (TODO, more wording? Use err?)",
		})
		return
	}

	meta = c
	return
}
