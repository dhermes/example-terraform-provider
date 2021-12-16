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
	"time"

	"github.com/google/uuid"

	"github.com/dhermes/example-terraform-provider/pkg/model"
)

// NOTE: Ensure that
//       * `addBook` satisfies `handleFunc`.
var (
	_ handleFunc = getAuthorByID
)

func addBook(w http.ResponseWriter, req *http.Request) {
	if notAllowed(w, req, http.MethodPost) {
		return
	}
	if contentTypeNotJSON(w, req) {
		return
	}
	if req.URL.Path != "/v1alpha1/book" {
		notFound(w)
		return
	}

	var abr addBookRequest
	if invalidJSONBody(w, req, &abr) {
		return
	}

	ctx := req.Context()
	pool := model.GetPool(ctx)
	b := model.Book{AuthorID: abr.AuthorID, Title: abr.Title, PublishDate: abr.PublishDate}
	id, err := model.InsertBook(ctx, pool, b)
	if err != nil {
		w.Header().Set(HeaderContentType, ContentTypeApplicationJSON)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "failed to insert book"}`+"\n")
		return
	}

	response := addBookResponse{BookID: id.String()}
	serializeJSONResponse(w, response)
}

type addBookRequest struct {
	Title       string     `json:"title"`
	AuthorID    uuid.UUID  `json:"author_id"`
	PublishDate *time.Time `json:"publish_date"`
}

type addBookResponse struct {
	BookID string `json:"book_id"`
}
