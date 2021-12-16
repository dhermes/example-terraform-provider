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
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/dhermes/example-terraform-provider/pkg/booksclient"
	"github.com/dhermes/example-terraform-provider/pkg/terraform"
)

func idFromString(idStr string) (uuid.UUID, error) {
	id, err := uuid.Parse(idStr)
	if err == nil {
		return id, nil
	}

	err = terraform.DiagnosticError{
		Summary: "Could not determine ID",
		Detail:  "Invalid ID parameter value",
	}
	return uuid.Nil, err
}

func authorFromResourceData(d *schema.ResourceData) (*booksclient.Author, error) {
	firstName, ok := d.Get("first_name").(string)
	if !ok {
		err := terraform.DiagnosticError{
			Summary: "Could not determine author first name",
			Detail:  "Invalid first name parameter type",
		}
		return nil, err
	}
	lastName, ok := d.Get("last_name").(string)
	if !ok {
		err := terraform.DiagnosticError{
			Summary: "Could not determine author last name",
			Detail:  "Invalid last name parameter type",
		}
		return nil, err
	}

	a := booksclient.Author{FirstName: firstName, LastName: lastName}
	idStr := d.Id()
	if idStr != "" {
		id, diags := idFromString(idStr)
		if diags != nil {
			return nil, diags
		}
		a.ID = &id
	}

	return &a, nil
}
