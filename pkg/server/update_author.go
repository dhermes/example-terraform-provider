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
//       * `updateAuthor` satisfies `handleFunc`.
var (
	_ handleFunc = updateAuthor
)

func updateAuthor(w http.ResponseWriter, req *http.Request) {
	if notAllowed(w, req, http.MethodPut) {
		return
	}
	if contentTypeNotJSON(w, req) {
		return
	}
	if req.URL.Path != "/v1alpha1/author" {
		notFound(w)
		return
	}

	var uar updateAuthorRequest
	if invalidJSONBody(w, req, &uar) {
		return
	}

	ctx := req.Context()
	pool := model.GetPool(ctx)
	a := model.Author{ID: uar.ID, FirstName: uar.FirstName, LastName: uar.LastName}
	err := model.UpdateAuthor(ctx, pool, a)
	if err != nil {
		w.Header().Set(HeaderContentType, ContentTypeApplicationJSON)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "failed to update author"}`+"\n")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

type updateAuthorRequest struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}
