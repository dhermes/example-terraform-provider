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
	"net"
	"net/http"
)

// Run registers all API routes and runs the Books API server.
func Run(ctx context.Context, c Config) error {
	m := http.NewServeMux()

	m.HandleFunc("/v1alpha1/author", addAuthor)
	m.HandleFunc("/v1alpha1/authors", getAuthors)
	m.HandleFunc("/v1alpha1/authors/", getAuthor)
	m.HandleFunc("/v1alpha1/book", addBook)
	m.HandleFunc("/v1alpha1/books", getBooks)
	m.HandleFunc("/v1alpha1/books/", getBook)
	m.HandleFunc("/", defaultHandler)

	s := &http.Server{
		Addr:        c.Addr,
		Handler:     m,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
	}
	return s.ListenAndServe()
}
