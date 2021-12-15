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

package server

import (
	"context"
	"database/sql"

	"github.com/dhermes/example-terraform-provider/pkg/model"
)

// Context uses `context.WithValue()` to attach configurated values that will
// be constant within the scope of the application.
func Context(ctx context.Context, c Config) (context.Context, error) {
	pool, err := sql.Open("pgx", c.DSN)
	if err != nil {
		return nil, err
	}

	err = pool.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	ctx = model.WithPool(ctx, pool)
	return ctx, nil
}
