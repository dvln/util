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

package dir

import "github.com/dvln/util/path"

// AbsPathify takes a path and attempts to clean it up and turn
// it into an absolute path via filepath.Clean and filepath.Abs
func AbsPathify(inPath string) string {
	return path.AbsPathify(inPath)
}

// Exists checks if given dir exists
func Exists(dir string) (bool, error) {
	return path.Exists(dir)
}

// CreateIfNotExists creates a file or a directory only if it does not already exist.
func CreateIfNotExists(dir string) error {
	return path.CreateIfNotExists(dir, true)
}
