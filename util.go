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

// Package util is an umbrella package for utility functions to augment the stdlib,
// these are small sub-packages that could be broken out into standalone eventually.
//   util/file - file focused utility routines
//   util/dir - directory focused utility routines
//   util/symlink - symlink focused utility routines
//   util/path - general path (file or dir) focused utility routines
//   util/system - routines for common system level examination
//   util/unit - unit conversion utility routines (to human format, from human format)
//   util/gotype - routines around manipulating/searching maps, slices, etc
// Right now these are independent packages, but all versioned within the single
// repo named 'github.com/dvln/util'
package util

