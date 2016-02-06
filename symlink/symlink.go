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

// package symlink is targeted towards reading symlinks with some intelligence
// towards making sure they are insuring they are pointed at the desired file
// or directory type
package symlink

import (
	"fmt"
	"os"
	"path/filepath"
)

// ReadSymlink will scan the given symlink and return what it points at
// which might be a file or directory, if any issues an error is returned.
// See: ReadSymlinkedDirectory() and ReadSymlinkedFile(), they can be used
// to verify the symlink target is of the expected type if that is needed.
func ReadSymlink(path string) (string, error) {
	var realPath string
	var err error
	if realPath, err = filepath.Abs(path); err != nil {
		return "", fmt.Errorf("unable to get absolute path for %s: %s", path, err)
	}
	if realPath, err = filepath.EvalSymlinks(realPath); err != nil {
		return "", fmt.Errorf("failed to canonicalise path for %s: %s", path, err)
	}
	return realPath, nil
}

// ReadSymlinkedDirectory returns the target directory of a symlink.
// The target of the symbolic link may not be a file.
func ReadSymlinkedDirectory(path string) (string, error) {
	realPath, err := ReadSymlink(path)
	if err != nil {
		return "", err
	}
	realPathInfo, err := os.Stat(realPath)
	if err != nil {
		return "", fmt.Errorf("failed to stat target '%s' of '%s': %s", realPath, path, err)
	}
	if !realPathInfo.Mode().IsDir() {
		return "", fmt.Errorf("canonical path points to a file '%s'", realPath)
	}
	return realPath, nil
}

// ReadSymlinkedFile returns the target file of a symlink.
// The target of the symbolic link may not be a directory.
func ReadSymlinkedFile(path string) (string, error) {
	realPath, err := ReadSymlink(path)
	if err != nil {
		return "", err
	}
	realPathInfo, err := os.Stat(realPath)
	if err != nil {
		return "", fmt.Errorf("failed to stat target '%s' of '%s': %s", realPath, path, err)
	}
	if !realPathInfo.Mode().IsRegular() {
		return "", fmt.Errorf("canonical path points does not point to a file '%s'", realPath)
	}
	return realPath, nil
}
