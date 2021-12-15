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
)

// NOTE: Ensure that
//       * `addBook` satisfies `handleFunc`.
var (
	_ handleFunc = getAuthor
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

	fmt.Println("TODO: addBook(", abr, ")")
}

type addBookRequest struct {
	Title       string     `json:"title"`
	AuthorID    uint64     `json:"author_id,string"`
	PublishDate *time.Time `json:"publish_date"`
}
