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

// Config provides the core set of (CLI) inputs needed to run the Books
// API server.
type Config struct {
	Addr string
}

// NewConfig returns a new `Config` with all relevant defaults provided and
// options for overriding.
func NewConfig(opts ...Option) (Config, error) {
	c := Config{}
	for _, opt := range opts {
		err := opt(&c)
		if err != nil {
			return Config{}, err
		}
	}
	return c, nil
}
