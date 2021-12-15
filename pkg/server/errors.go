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
	"encoding/json"
	"fmt"
	"net/http"
)

func notAllowed(w http.ResponseWriter, req *http.Request, method string) bool {
	if req.Method == method {
		return false
	}

	w.Header().Set(HeaderContentType, ContentTypeApplicationJSON)
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(w, `{"error": "not allowed"}`+"\n")
	return true
}

func contentTypeNotJSON(w http.ResponseWriter, req *http.Request) bool {
	if req.Header.Get(HeaderContentType) == ContentTypeApplicationJSON {
		return false
	}

	w.Header().Set(HeaderContentType, ContentTypeApplicationJSON)
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, `{"error": "JSON requests only"}`+"\n")
	return true
}

func notFound(w http.ResponseWriter) {
	w.Header().Set(HeaderContentType, ContentTypeApplicationJSON)
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, `{"error": "not found"}`+"\n")
}

func invalidJSONBody(w http.ResponseWriter, req *http.Request, v interface{}) bool {
	d := json.NewDecoder(req.Body)
	d.DisallowUnknownFields()

	err := d.Decode(v)
	if err == nil {
		return false
	}

	w.Header().Set(HeaderContentType, ContentTypeApplicationJSON)
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, `{"error": "invalid request body"}`+"\n")
	return true
}

func serializeJSONResponse(w http.ResponseWriter, v interface{}) bool {
	// NOTE: It'd be nice to just do `json.NewEncoder(w).Encode(v)`, but we
	//       can't do this because we need to write the status code header before
	//       writing the response body.
	responseBody, err := json.Marshal(v)
	if err != nil {
		w.Header().Set(HeaderContentType, ContentTypeApplicationJSON)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "could not serialize response"}`+"\n")
		return false
	}

	w.Header().Set(HeaderContentType, ContentTypeApplicationJSON)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(responseBody)+"\n")
	return true
}
