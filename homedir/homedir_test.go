package homedir

import "testing"

// See if we can get a user home directory
func TestHomeDirCheck(t *testing.T) {
	dir := UserHomeDir()
	if dir == "" {
		t.Fatal("Should have found a home dir, but was empty")
	}
}
