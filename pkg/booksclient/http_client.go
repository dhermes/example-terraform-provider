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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
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

// NewHTTPClient returns a new `HTTPClient` with all relevant defaults provided and
// options for overriding.
func NewHTTPClient(opts ...Option) (HTTPClient, error) {
	hc := HTTPClient{}
	for _, opt := range opts {
		err := opt(&hc)
		if err != nil {
			return HTTPClient{}, err
		}
	}
	return hc, nil
}

// RawClient returns a standard libary HTTP client associated with this client.
//
// NOTE: For now this is just a stub wrapper around `http.DefaultClient` but
//       it's provided here to make the code easier to test at a later date.
func (HTTPClient) RawClient() *http.Client {
	return http.DefaultClient
}

// AddAuthor adds a new author to be stored in the books service.
func (hc *HTTPClient) AddAuthor(ctx context.Context, a Author) (*AddAuthorResponse, error) {
	url := fmt.Sprintf("%s/v1alpha1/author", hc.Addr)
	asJSON, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(asJSON))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := hc.RawClient().Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("failed to add author (status %d, body %q)", resp.StatusCode, body)
	}

	var response AddAuthorResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateAuthor updates an author stored in the books service.
func (hc *HTTPClient) UpdateAuthor(ctx context.Context, a Author) (*Empty, error) {
	url := fmt.Sprintf("%s/v1alpha1/author", hc.Addr)
	asJSON, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewReader(asJSON))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := hc.RawClient().Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNoContent {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("failed to update author (status %d, body %q)", resp.StatusCode, body)
	}

	return &Empty{}, nil
}

// GetAuthorByID gets an author currently stored in the books service by ID.
func (hc *HTTPClient) GetAuthorByID(ctx context.Context, gabir GetAuthorByIDRequest) (*Author, error) {
	url := fmt.Sprintf("%s/v1alpha1/authors/%s", hc.Addr, url.PathEscape(gabir.AuthorID.String()))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := hc.RawClient().Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("failed to get author by ID (status %d, body %q)", resp.StatusCode, body)
	}

	var response Author
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetAuthorByName gets an author currently stored in the books service by name.
func (hc *HTTPClient) GetAuthorByName(ctx context.Context, gabnr GetAuthorByNameRequest) (*Author, error) {
	url := fmt.Sprintf("%s/v1alpha1/author", hc.Addr)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	q := req.URL.Query()
	q.Add("first_name", gabnr.FirstName)
	q.Add("last_name", gabnr.LastName)
	req.URL.RawQuery = q.Encode()

	resp, err := hc.RawClient().Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("failed to get author by name (status %d, body %q)", resp.StatusCode, body)
	}

	var response Author
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetAuthors gets all authors currently stored in the books service.
func (hc *HTTPClient) GetAuthors(ctx context.Context, _ Empty) (*GetAuthorsResponse, error) {
	url := fmt.Sprintf("%s/v1alpha1/authors", hc.Addr)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := hc.RawClient().Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("failed to get authors (status %d, body %q)", resp.StatusCode, body)
	}

	var response GetAuthorsResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteAuthorRequest deletes an author stored in the books service.
func (hc *HTTPClient) DeleteAuthorByID(ctx context.Context, dar DeleteAuthorRequest) (*Empty, error) {
	url := fmt.Sprintf("%s/v1alpha1/authors/%s", hc.Addr, url.PathEscape(dar.AuthorID.String()))
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := hc.RawClient().Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNoContent {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("failed to delete author by ID (status %d, body %q)", resp.StatusCode, body)
	}

	return &Empty{}, nil
}

// AddBook adds a new book to be stored in the books service.
//
// Before adding a book, a valid author must be created via `AddAuthor()`.
func (hc *HTTPClient) AddBook(ctx context.Context, b Book) (*AddBookResponse, error) {
	url := fmt.Sprintf("%s/v1alpha1/book", hc.Addr)
	asJSON, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(asJSON))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := hc.RawClient().Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("failed to add book (status %d, body %q)", resp.StatusCode, body)
	}

	var response AddBookResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateBook updates a book stored in the books service.
func (hc *HTTPClient) UpdateBook(ctx context.Context, b Book) (*Empty, error) {
	url := fmt.Sprintf("%s/v1alpha1/book", hc.Addr)
	asJSON, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewReader(asJSON))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := hc.RawClient().Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNoContent {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("failed to update book (status %d, body %q)", resp.StatusCode, body)
	}

	return &Empty{}, nil
}

// GetBookByID gets a book currently stored in the books service by ID.
func (hc *HTTPClient) GetBookByID(ctx context.Context, gbbir GetBookByIDRequest) (*Book, error) {
	url := fmt.Sprintf("%s/v1alpha1/books/%s", hc.Addr, url.PathEscape(gbbir.BookID.String()))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := hc.RawClient().Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("failed to get book by ID (status %d, body %q)", resp.StatusCode, body)
	}

	var response Book
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetBooks gets all books currently stored in the books service for a given author.
func (hc *HTTPClient) GetBooks(ctx context.Context, gbr GetBooksRequest) (*GetBooksResponse, error) {
	url := fmt.Sprintf("%s/v1alpha1/books", hc.Addr)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	q := req.URL.Query()
	q.Add("author_id", gbr.AuthorID.String())
	req.URL.RawQuery = q.Encode()

	resp, err := hc.RawClient().Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("failed to get books (status %d, body %q)", resp.StatusCode, body)
	}

	var response GetBooksResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteBookByID deletes a book stored in the books service.
func (hc *HTTPClient) DeleteBookByID(ctx context.Context, dbr DeleteBookRequest) (*Empty, error) {
	url := fmt.Sprintf("%s/v1alpha1/books/%s", hc.Addr, url.PathEscape(dbr.BookID.String()))
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := hc.RawClient().Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNoContent {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("failed to delete book by ID (status %d, body %q)", resp.StatusCode, body)
	}

	return &Empty{}, nil
}
