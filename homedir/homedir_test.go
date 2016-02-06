package homedir

import (
	"os"
	"testing"
	"runtime"
)

// See if we can get a user home directory
func TestHomeDirCheck(t *testing.T) {
	dir := UserHomeDir()
	if dir == "" {
		t.Fatal("Should have found a home dir, but was empty")
	}
	if runtime.GOOS != "windows" {
		// If not on Windows lets verify the Windows logic at least works...
		// will give better coverage numbers for this tiny package.  :)
		os.Setenv("TESTWINDOZE", "1")
		dir := UserHomeDir()
		if dir != "" {
			t.Fatalf("Should have had an empty fake windows home dir but contained: %s\n", dir)
		}
		os.Setenv("USERPROFILE", "C:\\some\\dir")
		dir = UserHomeDir()
		if dir == "" || dir != "C:\\some\\dir" {
			t.Fatal("Should have found a windows home dir set to \"C:\\some\\dir\", but found: \"%s\"\n", dir)
		}
		os.Setenv("HOMEDRIVE", "C:")
		os.Setenv("HOMEPATH", "\\some\\dir")
		dir = UserHomeDir()
		if dir == "" || dir != "C:\\some\\dir" {
			t.Fatal("Should have found a windows home dir set to \"C:\\some\\dir\", but found: \"%s\"\n", dir)
		}
	}
}
