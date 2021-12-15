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

package main

import (
	"fmt"
	"os"

	"github.com/dhermes/golembic/command"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/dhermes/example-terraform-provider/pkg/sqlmigrations"
)

// NOTE: Ensure that
//       * `sqlmigrations.AllMigrations` satisfies `command.RegisterMigrations`.
var (
	_ command.RegisterMigrations = sqlmigrations.AllMigrations
)

func run() error {
	cmd, err := command.MakeRootCommand(sqlmigrations.AllMigrations)
	if err != nil {
		return err
	}

	return cmd.Execute()
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
