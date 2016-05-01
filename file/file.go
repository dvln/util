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

// Package file contains some basic file existence, copying, cleanup,
// and matching functions.
package file

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/dvln/out"
	"github.com/dvln/util/path"
)

// Exists checks if given file exists, if you want to check for a directory
// use the dir.Exists() routine or if you want to check for both file and
// directory use the path.Exists() routine.
func Exists(file string) (bool, error) {
	exists, err := path.Exists(file)
	if err != nil {
		// error already wrapped by path.Exists()
		return exists, err
	}
	if exists {
		var fileinfo os.FileInfo
		fileinfo, err = os.Stat(file)
		if err != nil {
			return false, out.WrapErr(err, "Failed to stat file, unable to verify existence", 4013)
		}
		if fileinfo.IsDir() {
			exists = false
			err = out.NewErr("Item is a directory hence the file existence check failed", 4014)
		}
	}
	return exists, err
}

// CopyFile copies from src to dst until either EOF is reached
// on src or an error occurs. It verifies src exists and remove
// the dst if it exists.
func CopyFile(src, dst string) (int64, error) {
	cleanSrc := filepath.Clean(src)
	cleanDst := filepath.Clean(dst)
	if cleanSrc == cleanDst {
		return 0, nil
	}
	sf, err := os.Open(cleanSrc)
	if err != nil {
		return 0, out.WrapErr(err, "Failed to open source file", 4004)
	}
	defer sf.Close()
	if err := os.Remove(cleanDst); err != nil && !os.IsNotExist(err) {
		return 0, out.WrapErr(err, "Failed to clean/remove destination file", 4005)
	}
	df, err := os.Create(cleanDst)
	if err != nil {
		return 0, out.WrapErr(err, "Failed to create destination file", 4006)
	}
	defer df.Close()
	bytes, err := io.Copy(df, sf)
	if err != nil {
		return bytes, out.WrapErr(err, "Failed to copy source file to destination", 4007)
	}
	return bytes, nil
}

// CopyFileSetPerms copies from src to dst until either EOF is reached
// on src or an error occurs.  It will create the destination file with
// the specified permissions and return the number of bytes written (int64)
// and an error (nil if no error).
// Note: if destination file exists it will be removed by this routine
func CopyFileSetPerms(src, dst string, mode os.FileMode) (int64, error) {
	cleanSrc := filepath.Clean(src)
	cleanDst := filepath.Clean(dst)
	if cleanSrc == cleanDst {
		return 0, nil
	}
	sf, err := os.Open(cleanSrc)
	if err != nil {
		return 0, out.WrapErr(err, "Failed to open source file", 4004)
	}
	defer sf.Close()
	if err := os.Remove(cleanDst); err != nil && !os.IsNotExist(err) {
		return 0, out.WrapErr(err, "Failed to clean/remove destination file", 4005)
	}
	oldUmask := syscall.Umask(0)
	defer syscall.Umask(oldUmask)
	df, err := os.OpenFile(cleanDst, syscall.O_CREAT|syscall.O_EXCL|syscall.O_APPEND|syscall.O_WRONLY, mode)
	if err != nil {
		return 0, out.WrapErr(err, "Failed to create destination file", 4006)
	}
	defer df.Close()
	bytes, err := io.Copy(df, sf)
	if err != nil {
		return bytes, out.WrapErr(err, "Failed to copy source file to destination", 4007)
	}
	return bytes, nil
}

// CreateIfNotExists creates a file or a directory only if it does not already exist.
func CreateIfNotExists(file string) error {
	return path.CreateIfNotExists(file, false)
}

// exclusion return true if the specified pattern is an exclusion
func exclusion(pattern string) bool {
	return pattern[0] == '!'
}

// empty return true if the specified pattern is empty
func empty(pattern string) bool {
	return pattern == ""
}

// CleanPatterns takes a slice of patterns returns a new
// slice of patterns cleaned with filepath.Clean, stripped
// of any empty patterns and lets the caller know whether the
// slice contains any exception patterns (prefixed with !).
func CleanPatterns(patterns []string) ([]string, [][]string, bool, error) {
	// Loop over exclusion patterns and:
	// 1. Clean them up.
	// 2. Indicate whether we are dealing with any exception rules.
	// 3. Error if we see a single exclusion marker on it's own (!).
	cleanedPatterns := []string{}
	patternDirs := [][]string{}
	exceptions := false
	for _, pattern := range patterns {
		// Eliminate leading and trailing whitespace.
		pattern = strings.TrimSpace(pattern)
		if empty(pattern) {
			continue
		}
		if exclusion(pattern) {
			if len(pattern) == 1 {
				return nil, nil, false, out.NewErr("Illegal exclusion pattern: !", 4009)
			}
			exceptions = true
		}
		pattern = filepath.Clean(pattern)
		cleanedPatterns = append(cleanedPatterns, pattern)
		if exclusion(pattern) {
			pattern = pattern[1:]
		}
		patternDirs = append(patternDirs, strings.Split(pattern, "/"))
	}

	return cleanedPatterns, patternDirs, exceptions, nil
}

// Matches returns true if file matches any of the patterns
// and isn't excluded by any of the subsequent patterns.
func Matches(file string, patterns []string) (bool, error) {
	file = filepath.Clean(file)

	if file == "." {
		// Don't let them exclude everything, kind of silly.
		return false, nil
	}

	patterns, patDirs, _, err := CleanPatterns(patterns)
	if err != nil {
		return false, out.WrapErr(err, "Unable to clean all patterns", 4010)
	}

	return OptimizedMatches(file, patterns, patDirs)
}

// OptimizedMatches is basically the same as fileutils.Matches() but optimized for archive.go.
// It will assume that the inputs have been preprocessed and therefore the function
// doen't need to do as much error checking and clean-up. This was done to avoid
// repeating these steps on each file being checked during the archive process.
// The more generic fileutils.Matches() can't make these assumptions.
func OptimizedMatches(file string, patterns []string, patDirs [][]string) (bool, error) {
	matched := false
	parentPath := filepath.Dir(file)
	parentPathDirs := strings.Split(parentPath, "/")

	for i, pattern := range patterns {
		negative := false

		if exclusion(pattern) {
			negative = true
			pattern = pattern[1:]
		}

		match, err := filepath.Match(pattern, file)
		if err != nil {
			return false, out.WrapErr(err, "Optimized matching, failed to match pattern", 4008)
		}

		if !match && parentPath != "." {
			// Check to see if the pattern matches one of our parent dirs.
			if len(patDirs[i]) <= len(parentPathDirs) {
				match, _ = filepath.Match(strings.Join(patDirs[i], "/"),
					strings.Join(parentPathDirs[:len(patDirs[i])], "/"))
			}
		}

		if match {
			matched = !negative
		}
	}

	if matched {
		out.Debugf("Skipping excluded path: %s", file)
	}
	return matched, nil
}

