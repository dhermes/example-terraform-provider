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
	"strings"

	"github.com/google/uuid"

	"github.com/dhermes/example-terraform-provider/pkg/model"
)

// NOTE: Ensure that
//       * `getAuthor` satisfies `handleFunc`.
var (
	_ handleFunc = getAuthorByID
)

func getAuthorByID(w http.ResponseWriter, req *http.Request) {
	if notAllowed(w, req, http.MethodGet) {
		return
	}
	if contentTypeNotJSON(w, req) {
		return
	}
	if !strings.HasPrefix(req.URL.Path, "/v1alpha1/authors/") {
		notFound(w)
		return
	}

	suffix := strings.TrimPrefix(req.URL.Path, "/v1alpha1/authors/")
	id, err := uuid.Parse(suffix)
	if err != nil {
		w.Header().Set(HeaderContentType, ContentTypeApplicationJSON)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "invalid ID"}`+"\n")
		return
	}

	ctx := req.Context()
	pool := model.GetPool(ctx)
	a, err := model.GetAuthorByID(ctx, pool, id)
	if err != nil {
		w.Header().Set(HeaderContentType, ContentTypeApplicationJSON)
		// TODO: Consider supporting a 404 here
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "failed to get author by ID"}`+"\n")
		return
	}

	serializeJSONResponse(w, dbAuthorToResult(a))
}

type authorResponse struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BookCount uint32 `json:"book_count"`
}

func dbAuthorToResult(a *model.Author) authorResponse {
	return authorResponse{
		ID:        a.ID.String(),
		FirstName: a.FirstName,
		LastName:  a.LastName,
		BookCount: a.BookCount,
	}
}
