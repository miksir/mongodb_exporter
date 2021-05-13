// Copyright 2017 Percona LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common

import "strings"

// IsSystemDB tests whether system db name passed.
func IsSystemDB(dbName string) bool {
	switch dbName {
	case "admin", "config", "local", "mdb_internal":
		return true
	default:
		return false
	}
}

// IsSystemCollection tests whether system collection name passed.
func IsSystemCollection(collName string) bool {
	return strings.HasPrefix(collName, "system.")
}
