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
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/dhermes/example-terraform-provider/pkg/terraform"
)

func idFromString(idStr, fieldName string) (uuid.UUID, error) {
	id, err := uuid.Parse(idStr)
	if err == nil {
		return id, nil
	}

	err = terraform.DiagnosticError{
		Summary: fmt.Sprintf("Could not determine %s", fieldName),
		Detail:  "Invalid UUID parameter value",
	}
	return uuid.Nil, err
}

func dateFromString(dateStr, fieldName string) (Date, error) {
	t, err := time.ParseInLocation(dateLayout, dateStr, time.UTC)
	if err == nil {
		return Date{Time: t}, nil
	}

	err = terraform.DiagnosticError{
		Summary: fmt.Sprintf("Could not determine %s", fieldName),
		Detail:  "Invalid date parameter value",
	}
	return Date{}, err
}
