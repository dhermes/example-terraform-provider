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
				DefaultFunc: schema.EnvDefaultFunc(booksprovider.EnvVarBooksAPIAddr, nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"books_api_author": resourceAuthor(),
			"books_api_book":   resourceBook(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"books_api_author": dataSourceAuthor(),
		},
		ConfigureContextFunc: configureContext,
	}
}

func configureContext(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	c, err := booksprovider.Configure(ctx, d)
	if err == nil {
		return c, nil
	}

	return nil, terraform.AppendDiagnostic(err, nil)
}
