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
