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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAuthor returns the `author` resource in the Terraform provider for
// the Books API.
func resourceAuthor() *schema.Resource {
	return &schema.Resource{
		CreateContext: nil, // TODO
		ReadContext:   nil, // TODO
		UpdateContext: nil, // TODO
		DeleteContext: nil, // TODO
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
			// TODO
		},
	}
}