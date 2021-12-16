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
//       * `getAuthorByName` satisfies `handleFunc`.
var (
	_ handleFunc = getAuthorByName
)

func getAuthorByName(w http.ResponseWriter, req *http.Request) {
	if notAllowed(w, req, http.MethodGet) {
		return
	}
	if contentTypeNotJSON(w, req) {
		return
	}
	if req.URL.Path != "/v1alpha1/author" {
		notFound(w)
		return
	}

	// NOTE: We could be much more restrictive here with known / unkown inputs.
	q := req.URL.Query()
	firstName := q.Get("first_name")
	lastName := q.Get("last_name")
	if firstName == "" || lastName == "" {
		w.Header().Set(HeaderContentType, ContentTypeApplicationJSON)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "missing author first or last name query parameter"}`+"\n")
		return
	}

	ctx := req.Context()
	pool := model.GetPool(ctx)
	a, err := model.GetAuthorByName(ctx, pool, firstName, lastName)
	if err != nil {
		w.Header().Set(HeaderContentType, ContentTypeApplicationJSON)
		// TODO: Consider supporting a 404 here
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "failed to get author by name"}`+"\n")
		return
	}

	serializeJSONResponse(w, dbAuthorToResult(a))
}
