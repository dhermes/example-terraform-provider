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

package booksprovider

import (
	"encoding/json"
	"fmt"
	"time"
)

// NOTE: Ensure that
//       * `Date` satisfies `json.Unmarshaler`.
//       * `Date` satisfies `json.Marshaler`.
//       * `Date` satisfies `fmt.Stringer`.
var (
	_ json.Unmarshaler = (*Date)(nil)
	_ json.Marshaler   = Date{}
	_ fmt.Stringer     = Date{}
)

const (
	dateLayout = "2006-01-02"
)

type Date struct {
	time.Time
}

// UnmarshalJSON unmarshals a string in `YYYY-MM-DD` format into a `Date`.
func (d *Date) UnmarshalJSON(data []byte) error {
	s := ""
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	t, err := time.ParseInLocation(dateLayout, s, time.UTC)
	if err != nil {
		return err
	}

	d.Time = t
	return nil
}

// MarshalJSON marshals a `Date` into a JSON string in `YYYY-MM-DD` format.
func (d Date) MarshalJSON() ([]byte, error) {
	s := d.String()
	return json.Marshal(s)
}

// String presents the `Date` as a string in `YYYY-MM-DD` format.
func (d Date) String() string {
	return d.Time.Format(dateLayout)
}
