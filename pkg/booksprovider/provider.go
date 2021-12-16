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
	"net/url"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/dhermes/example-terraform-provider/pkg/booksclient"
)

// ConfigureContext is a function for configuring the provider.
func ConfigureContext(_ context.Context, d *schema.ResourceData) (booksclient.Client, error) {
	addr, ok := d.Get("addr").(string)
	if addr == "" || !ok {
		err := DiagnosticError{
			Summary: "Unable to create Books API client",
			Detail:  "Unable to determine address for Books API",
		}
		return nil, err
	}

	_, err := url.Parse(addr)
	if err != nil {
		err := DiagnosticError{
			Summary: "Unable to create Books API client",
			Detail:  "Failed to parse Books API base address as a URL",
		}
		return nil, err
	}

	c, err := booksclient.NewHTTPClient(booksclient.OptAddr(addr))
	if err != nil {
		return nil, err
	}

	return &c, nil
}
