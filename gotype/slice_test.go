// These are various utility routines from docker, viper and various other
// tools/packages along with any local additions/mods.  For any local mods
// the Apache license is included:
//
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

package gotype

import "testing"

// See if string inside slice check is working
func TestStringInSlice(t *testing.T) {
	test := []string{"hello", "bye", "yourmama"}
	inside := StringInSlice("goodbye", test)
	if inside {
		t.Fatal("Found string in slice that should not have had the string inside of it")
	}

	inside = StringInSlice("bye", test)
	if !inside {
		t.Fatal("Failed to find string in slice, should have been there")
	}
}
