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

package sqlmigrations

import (
	"github.com/dhermes/golembic"
)

// AllMigrations returns a sequence of migrations.
func AllMigrations() (*golembic.Migrations, error) {
	root, err := golembic.NewMigration(
		golembic.OptRevision("332b84ad0543"),
		golembic.OptDescription("Create authors table"),
		golembic.OptUp(AddAuthorsTable),
	)
	if err != nil {
		return nil, err
	}

	migrations, err := golembic.NewSequence(*root)
	if err != nil {
		return nil, err
	}

	err = migrations.RegisterManyOpt(
		[]golembic.MigrationOption{
			golembic.OptPrevious("332b84ad0543"),
			golembic.OptRevision("8fc64f953bb0"),
			golembic.OptDescription("Create books table"),
			golembic.OptUp(AddBooksTable),
		},
	)
	if err != nil {
		return nil, err
	}

	return migrations, nil

}
