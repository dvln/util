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

package url

import (
	"testing"
)

// TestGetScheme provides a simple coverage test to see if we can
// properly figure out the URL "scheme"
func TestGetScheme(t *testing.T) {
	url1 := "http://github.com/dvln/dvln"
	scheme := GetScheme(url1)
	if scheme != "http" {
		t.Fatalf("Failed to parse URL 1 (%s) correctly, scheme should have been \"http\", was: %s", url1, scheme)
	}

	url2 := "github.com/dvln/dvln"
	scheme = GetScheme(url2)
	if scheme != "" {
		t.Fatalf("Failed to parse URL 2 (%s) correctly, scheme should have been \"\", was: %s", url1, scheme)
	}
}
