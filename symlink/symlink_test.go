package symlink

import (
	"os"
	"path/filepath"
	"testing"
)

// Reading a symlink to a directory must return the directory
func TestReadSymlinkedDirectoryExistingDirectory(t *testing.T) {
	var err error

	// defer removal of test dirs/links in case we fail early for any reason,
	// note that we do removal with checks below before we call these defers
	// so it's very likely the dirs are already gone (if not we'll have err'd,
	// but if we error before those removes lets clean em up anyhow for now)
	tmpDir := os.TempDir()
	dir := filepath.Join(tmpDir, "testReadSymlinkToExistingDirectory")
	symlink := filepath.Join(tmpDir, "dirLinkTest")
	defer os.Remove(dir)
	defer os.Remove(symlink)

	if err = os.Mkdir(dir, 0777); err != nil {
		t.Errorf("failed to create directory: %s", err)
	}

	if err = os.Symlink(dir, symlink); err != nil {
		t.Errorf("failed to create symlink: %s", err)
	}

	var path string
	if path, err = ReadSymlinkedDirectory(symlink); err != nil {
		t.Fatalf("failed to read symlink to directory: %s", err)
	}
	if !(path == dir ||
		path == filepath.Join("/private", dir)) {
		t.Fatalf("symlink returned unexpected directory: %s", path)
	}

	if path, err = ReadSymlinkedFile(symlink); err == nil {
		t.Fatalf("ReadSymlinkedFile on a symlink to a dir should not have worked")
	}
	if path != "" {
		t.Fatalf("path should be empty: %s", path)
	}

	if path, err = ReadSymlink(symlink); err != nil {
		t.Fatalf("ReadSymlink on a symlink to a anything should've worked")
	}
	if path == "" {
		t.Fatalf("path should not be empty: %s", path)
	}

	if err = os.Remove(dir); err != nil {
		t.Errorf("failed to remove temporary directory: %s", err)
	}
	if err = os.Remove(symlink); err != nil {
		t.Errorf("failed to remove symlink: %s", err)
	}
}

// Reading a non-existing symlink must fail
func TestReadSymlinkedDirectoryNonExistingSymlink(t *testing.T) {
	var path string
	var err error
	tmpDir := os.TempDir()
	nonExistentPath := filepath.Join(tmpDir, "NonExistentPath")

	if path, err = ReadSymlinkedDirectory(nonExistentPath); err == nil {
		t.Fatalf("error expected for non-existing symlink")
	}
	if path != "" {
		t.Fatalf("expected empty path, but '%s' was returned", path)
	}

	if path, err = ReadSymlinkedFile(nonExistentPath); err == nil {
		t.Fatalf("error expected for non-existing symlink")
	}
	if path != "" {
		t.Fatalf("expected empty path, but '%s' was returned", path)
	}

	if path, err = ReadSymlink(nonExistentPath); err == nil {
		t.Fatalf("error expected for non-existing symlink")
	}
	if path != "" {
		t.Fatalf("expected empty path, but '%s' was returned", path)
	}
}

// Reading a symlink to a file must fail
func TestReadSymlinkedDirectoryToFile(t *testing.T) {
	var err error
	var file *os.File

	tmpDir := os.TempDir()
	filename := filepath.Join(tmpDir, "testReadSymlinkToFile")
	symlink := filepath.Join(tmpDir, "fileLinkTest")
	defer os.Remove(filename)
	defer os.Remove(symlink)
	if file, err = os.Create(filename); err != nil {
		t.Fatalf("failed to create file: %s", err)
	}
	file.Close()

	if err = os.Symlink(filename, symlink); err != nil {
		t.Errorf("failed to create symlink: %s", err)
	}

	var path string
	if path, err = ReadSymlinkedDirectory(symlink); err == nil {
		t.Fatalf("ReadSymlinkedDirectory on a symlink to a file should've failed")
	}
	if path != "" {
		t.Fatalf("path should've been empty: %s", path)
	}

	if path, err = ReadSymlinkedFile(symlink); err != nil {
		t.Fatalf("ReadSymlinkedFile on a symlink to a file should've worked")
	}
	if path == "" {
		t.Fatalf("path should not be empty: %s", path)
	}

	if path, err = ReadSymlink(symlink); err != nil {
		t.Fatalf("ReadSymlink on a symlink to a anything should've worked")
	}
	if path == "" {
		t.Fatalf("path should not be empty: %s", path)
	}

	if err = os.Remove(filename); err != nil {
		t.Errorf("failed to remove file: %s", err)
	}
	if err = os.Remove(symlink); err != nil {
		t.Errorf("failed to remove symlink: %s", err)
	}
}
