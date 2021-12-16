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

	"github.com/dhermes/example-terraform-provider/pkg/model"
)

// NOTE: Ensure that
//       * `oneAuthorDispatch` satisfies `handleFunc`.
//       * `addAuthor` satisfies `handleFunc`.
var (
	_ handleFunc = oneAuthorDispatch
	_ handleFunc = addAuthor
)

// oneAuthorDispatch dispatches to `addAuthor()` for a `POST` request and
// to `getAuthorByName()` for a `GET` request.
func oneAuthorDispatch(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		addAuthor(w, req)
		return
	}

	if req.Method == http.MethodGet {
		getAuthorByName(w, req)
		return
	}

	notAllowed(w, req, req.Method+"-mismatch") // Ensure the method is wrong
}

func addAuthor(w http.ResponseWriter, req *http.Request) {
	if notAllowed(w, req, http.MethodPost) {
		return
	}
	if contentTypeNotJSON(w, req) {
		return
	}
	if req.URL.Path != "/v1alpha1/author" {
		notFound(w)
		return
	}

	var aar addAuthorRequest
	if invalidJSONBody(w, req, &aar) {
		return
	}

	ctx := req.Context()
	pool := model.GetPool(ctx)
	a := model.Author{FirstName: aar.FirstName, LastName: aar.LastName}
	id, err := model.InsertAuthor(ctx, pool, a)
	if err != nil {
		w.Header().Set(HeaderContentType, ContentTypeApplicationJSON)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "failed to insert author"}`+"\n")
		return
	}

	response := addAuthorResponse{AuthorID: id.String()}
	serializeJSONResponse(w, response)
}

type addAuthorRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type addAuthorResponse struct {
	AuthorID string `json:"author_id"`
}
