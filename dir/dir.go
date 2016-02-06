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

// Package dir contains a few simple directory existence, cleanup,
// scanning and creation functions.
package dir

import (
	"os"
	"path/filepath"

	"github.com/dvln/out"
	"github.com/dvln/util/path"
)

// AbsPathify takes a path and attempts to clean it up and turn
// it into an absolute path via filepath.Clean and filepath.Abs
func AbsPathify(inPath string) string {
	return path.AbsPathify(inPath)
}

// Exists checks if given dir exists, if you want to check for a *file*
// use the file.Exists() routine or if you want to check for both file and
// directory use the path.Exists() routine.
func Exists(dir string) (bool, error) {
	exists, err := path.Exists(dir)
	if err != nil {
		// error already wrapped by path.Exists()
		return exists, err
	}
	if exists {
		fileinfo, err := os.Stat(dir)
		if err != nil {
			return false, out.WrapErr(err, "Failed to stat directory, unable to verify existence", 4011)
		}
		if !fileinfo.IsDir() {
			exists = false
			err = out.NewErr("Item is not a directory hence directory existence check failed", 4012)
		}
	}
	return exists, err
}

// CreateIfNotExists creates a file or a directory only if it does not already exist.
func CreateIfNotExists(dir string) error {
	return path.CreateIfNotExists(dir, true)
}

// FindDirInOrAbove will look for a given directory in or above the given
// starting dir (iit will travese "up" the filesystem and examine parent
// directories to see if they contain the given directory).  If the findDir
// dir is found then the dir it's found in will be returned, else "" (any
// unexpected error will come back in the error return parameter)
func FindDirInOrAbove(startDir string, findDir string) (string, error) {
	fullPath := filepath.Join(startDir, findDir)
	exists, err := Exists(fullPath)
	if err != nil {
		return "", out.WrapErr(err, "Problem checking directory existence", 4003)
	}
	if exists {
		return startDir, nil
	}
	baseDir := filepath.Dir(startDir)
	if baseDir == "." || (len(baseDir) == 1 && baseDir[0] == filepath.Separator) {
		return "", nil
	}
	return FindDirInOrAbove(baseDir, findDir)
}
