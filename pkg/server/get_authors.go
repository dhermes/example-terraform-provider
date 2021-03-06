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
//       * `getAuthors` satisfies `handleFunc`.
var (
	_ handleFunc = getAuthorByID
)

func getAuthors(w http.ResponseWriter, req *http.Request) {
	if notAllowed(w, req, http.MethodGet) {
		return
	}
	if contentTypeNotJSON(w, req) {
		return
	}
	if req.URL.Path != "/v1alpha1/authors" {
		notFound(w)
		return
	}

	ctx := req.Context()
	pool := model.GetPool(ctx)
	authorsDB, err := model.GetAllAuthors(ctx, pool)
	if err != nil {
		w.Header().Set(HeaderContentType, ContentTypeApplicationJSON)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "failed to get all authors"}`+"\n")
		return
	}

	authors := make([]authorResponse, len(authorsDB))
	for i, a := range authorsDB {
		authors[i] = dbAuthorToResult(&a)
	}
	response := authorsResponse{Authors: authors}
	serializeJSONResponse(w, response)
}

type authorsResponse struct {
	Authors []authorResponse `json:"authors"`
}
