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
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/dhermes/example-terraform-provider/pkg/booksclient"
)

func getClientFromMeta(meta interface{}) (booksclient.Client, diag.Diagnostics) {
	// meta: expected to be the output of `providerConfigure()`
	c, ok := meta.(booksclient.Client)
	if ok {
		return c, nil
	}

	diags := []diag.Diagnostic{
		{
			Severity: diag.Error,
			Summary:  "No Books API client available",
			Detail:   "Expected provider configure function to return a Books API client",
		},
	}
	return nil, diags
}

func idFromString(idStr string) (uuid.UUID, diag.Diagnostics) {
	id, err := uuid.Parse(idStr)
	if err == nil {
		return id, nil
	}

	diags := []diag.Diagnostic{
		{
			Severity: diag.Error,
			Summary:  "Could not determine ID",
			Detail:   "Invalid ID parameter value",
		},
	}
	return uuid.Nil, diags
}
