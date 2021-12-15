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
	"context"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/spf13/cobra"

	"github.com/dhermes/example-terraform-provider/pkg/server"
)

func run() error {
	ctx := context.Background()

	c, err := server.NewConfig()
	if err != nil {
		return err
	}
	cmd := &cobra.Command{
		Use:           "server",
		Short:         "Run Books API server",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(_ *cobra.Command, _ []string) error {
			ctx, err := server.Context(ctx, c)
			if err != nil {
				return err
			}

			return server.Run(ctx, c)
		},
	}

	cmd.PersistentFlags().StringVar(
		&c.Addr,
		"addr",
		c.Addr,
		"The bind address to use for the server, e.g. ':7534'",
	)
	cmd.PersistentFlags().StringVar(
		&c.DSN,
		"dsn",
		c.DSN,
		"The DSN to use for the database, e.g. 'postgres://books_app:testpassword_app@127.0.0.1:22411/books'",
	)

	required := []string{"addr", "dsn"}
	for _, name := range required {
		err := cobra.MarkFlagRequired(cmd.PersistentFlags(), name)
		if err != nil {
			return err
		}
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
