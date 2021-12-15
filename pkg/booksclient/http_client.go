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

package booksclient

import (
	"context"
	"errors"
)

// NOTE: Ensure that
//       * `HTTPClient` satisfies `Client`.
var (
	_ Client = (*HTTPClient)(nil)
)

// HTTPClient is an HTTP implementation for the Books API.
type HTTPClient struct {
	// Addr is the base address for the Books API, e.g. `http://localhost:7104`.
	Addr string
}

func (*HTTPClient) AddAuthor(_ context.Context, _ Author) (*AddAuthorResponse, error) {
	return nil, errors.New("not implemented: add author")
}

func (*HTTPClient) GetAuthors(_ context.Context, _ Empty) (*GetAuthorsResponse, error) {
	return nil, errors.New("not implemented: get authors")
}

func (*HTTPClient) AddBook(_ context.Context, _ Book) (*AddBookResponse, error) {
	return nil, errors.New("not implemented: add book")
}

func (*HTTPClient) GetBooks(_ context.Context, _ GetBooksRequest) (*GetBooksResponse, error) {
	return nil, errors.New("not implemented: get books")
}
