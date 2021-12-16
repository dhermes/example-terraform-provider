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
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/dhermes/example-terraform-provider/pkg/model"
)

// NOTE: Ensure that
//       * `getBooks` satisfies `handleFunc`.
var (
	_ handleFunc = getAuthorByID
)

func getBooks(w http.ResponseWriter, req *http.Request) {
	if notAllowed(w, req, http.MethodGet) {
		return
	}
	if contentTypeNotJSON(w, req) {
		return
	}
	if req.URL.Path != "/v1alpha1/books" {
		notFound(w)
		return
	}

	// NOTE: We could be much more restrictive here with known / unkown inputs.
	q := req.URL.Query()
	idStr := q.Get("author_id")
	if idStr == "" {
		w.Header().Set(HeaderContentType, ContentTypeApplicationJSON)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "missing author ID query parameter"}`+"\n")
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		w.Header().Set(HeaderContentType, ContentTypeApplicationJSON)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "invalid ID"}`+"\n")
		return
	}

	ctx := req.Context()
	pool := model.GetPool(ctx)
	booksDB, err := model.GetAllBooksByAuthor(ctx, pool, id)
	if err != nil {
		w.Header().Set(HeaderContentType, ContentTypeApplicationJSON)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "failed to get all books by author"}`+"\n")
		return
	}

	books := make([]bookResponse, len(booksDB))
	for i, b := range booksDB {
		books[i] = dbBookToResult(&b)
	}
	response := booksResponse{Books: books}
	serializeJSONResponse(w, response)

}

type booksResponse struct {
	Books []bookResponse `json:"books"`
}
