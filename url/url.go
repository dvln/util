// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package url (github.com/dvln/util/url) is for basic URL parsing
// or manipulation helper functions.
package url

import (
	"strings"
)

// GetScheme tries to parse the given URL (string) to see if it has
// a scheme (eg: in "https://github.com/dvln/dvln" the scheme would
// be "https", in "github.com/dvln/dvln" there is no scheme).  If
// there is no scheme then "" is returned, otherwise the scheme in
// use is returned.
func GetScheme(url string) string {
	parts := strings.Split(url, "://")
	if len(parts) >= 2 {
		return parts[0]
	}
	return ""
}
