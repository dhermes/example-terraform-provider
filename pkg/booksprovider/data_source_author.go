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
	"sync"

	"github.com/dhermes/example-terraform-provider/pkg/booksclient"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceAuthor represents a `books_api_author` data source and directly
// interacts with the Terraform plugin SDK to check for new or updated
// values.
type DataSourceAuthor struct {
	mutex sync.RWMutex
	d     *schema.ResourceData

	FirstName *string    `terraform:"first_name,required"`
	LastName  *string    `terraform:"last_name,required"`
	BookCount *int       `terraform:"book_count,computed"`
	ID        *uuid.UUID `terraform:"id,string,computed"`
}

// Read is the read (R) component of the CRUD lifecycle for
// the `books_api_author` data source.
//
// NOTE: This assumes the called has already invoked `dsa.Populate()`, either
//       directly or indirectly, e.g. via `NewDataSourceAuthor()`.
func (dsa *DataSourceAuthor) Read(ctx context.Context, c booksclient.Client) error {
	gabnr := booksclient.GetAuthorByNameRequest{FirstName: dsa.GetFirstName(), LastName: dsa.GetLastName()}
	a, err := c.GetAuthorByName(ctx, gabnr)
	if err != nil {
		return err
	}

	dsa.FirstName = &a.FirstName
	dsa.LastName = &a.LastName
	bc := int(a.BookCount)
	dsa.BookCount = &bc
	return dsa.Persist()
}
