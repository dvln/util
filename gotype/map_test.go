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

// See if turning map to lower case only keys is working
func TestInsensitiseMap(t *testing.T) {
	testMap := map[string]interface{}{
		"Hello": "value",
		"yourMama": "value",
		"TESTING": "value",
		"fun": "Value",
	}
	checkMap := map[string]string{
		"hello": "value",
		"yourmama": "value",
		"testing": "value",
		"fun": "Value",
	}
	InsensitiviseMap(testMap)
	// At this point testMap keys should all be lower case and match checkMap,
	// the values should be unchanged and also map checkMap (which as the same
	// values as the original testMap, or at least should)
	for key, val := range checkMap {
		testVal, ok := testMap[key]
		if !ok {
			t.Fatalf("Expected to find lower case key \"%s\" in results map, did not!\nMap:%+v\n", key, testMap)
		}
		if testVal.(string) != val {
			t.Fatalf("Map values should remain unchanged, key \"%s\" has orig val \"%s\" and final val \"%s\"\n", key, testVal.(string), val)
		}
	}
}
