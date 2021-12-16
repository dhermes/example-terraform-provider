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

package terraform

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

// NOTE: Ensure that
//       * `DiagnosticError` satisfies `error`.
//       * `DiagnosticsProvider` satisfies `error`.
var (
	_ error               = DiagnosticError{}
	_ DiagnosticsProvider = DiagnosticError{}
)

// DiagnosticsProvider is an extended error interface.
type DiagnosticsProvider interface {
	AppendDiagnostic(diag.Diagnostics) diag.Diagnostics
}

// DiagnosticError is an idiomatic Go error that seeks to match a subset of
// the behavior of Hashicorp `diag.Diagnostic`.
type DiagnosticError struct {
	Summary string
	Detail  string
}

// Error satisfies the `error` interface.
func (de DiagnosticError) Error() string {
	return de.Summary
}

// AppendDiagnostic appends a `diag.Diagnostic` from the current error.
func (de DiagnosticError) AppendDiagnostic(diags diag.Diagnostics) diag.Diagnostics {
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  de.Summary,
		Detail:   de.Detail,
	})
	return diags
}
