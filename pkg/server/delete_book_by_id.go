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
//       * `deleteBookByID` satisfies `handleFunc`.
var (
	_ handleFunc = deleteBookByID
)

func deleteBookByID(w http.ResponseWriter, req *http.Request) {
	if notAllowed(w, req, http.MethodDelete) {
		return
	}
	if !strings.HasPrefix(req.URL.Path, "/v1alpha1/books/") {
		notFound(w)
		return
	}

	suffix := strings.TrimPrefix(req.URL.Path, "/v1alpha1/books/")
	id, err := uuid.Parse(suffix)
	if err != nil {
		w.Header().Set(HeaderContentType, ContentTypeApplicationJSON)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "invalid ID"}`+"\n")
		return
	}

	ctx := req.Context()
	pool := model.GetPool(ctx)
	err = model.DeleteBookByID(ctx, pool, id)
	if err != nil {
		w.Header().Set(HeaderContentType, ContentTypeApplicationJSON)
		// TODO: Consider supporting a 404 here
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "failed to delete book by ID"}`+"\n")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
