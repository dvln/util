package path

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// Poor test on pathify, need some symlinks/etc eventually
func TestAbsPathify(t *testing.T) {
	tempFolder, err := ioutil.TempDir("", "dvln-util-path-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempFolder)

	results := AbsPathify(tempFolder)
	if results == "" {
		t.Fatal("AbsPathify() seems to have returned nothing unexpectedly")
	}
	if results != tempFolder {
		t.Fatalf("AbsPathify() messed up a sane path (%v): %v", tempFolder, results)
	}
}

func TestCreateIfNotExistsDir(t *testing.T) {
	tempFolder, err := ioutil.TempDir("", "dvln-util-path-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempFolder)

	folderToCreate := filepath.Join(tempFolder, "tocreate")

	if err := CreateIfNotExists(folderToCreate, true); err != nil {
		t.Fatal(err)
	}
	fileinfo, err := os.Stat(folderToCreate)
	if err != nil {
		t.Fatalf("Should have create a folder, got %v", err)
	}

	if !fileinfo.IsDir() {
		t.Fatalf("Should have been a dir, seems it's not")
	}
}

func TestCreateIfNotExistsFile(t *testing.T) {
	tempFolder, err := ioutil.TempDir("", "dvln-util-path-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempFolder)

	fileToCreate := filepath.Join(tempFolder, "file/to/create")

	if err := CreateIfNotExists(fileToCreate, false); err != nil {
		t.Fatal(err)
	}
	fileinfo, err := os.Stat(fileToCreate)
	if err != nil {
		t.Fatalf("Should have create a file, got %v", err)
	}

	if fileinfo.IsDir() {
		t.Fatalf("Should have been a file, seems it's not")
	}

	exists, err := Exists(fileToCreate)
	if err != nil {
		t.Fatalf("File should have been found & existed, got %v", err)
	}
	if !exists {
		t.Fatalf("Exists() failed to detect existence of file")
	}
}

func TestExistsIfNotExists(t *testing.T) {
	exists, err := Exists(".bogusfile")
	if err != nil {
		t.Fatal("Missing file should return false, but err should be nil")
	}
	if exists {
		t.Fatal("Bogus file should not have existed")
	}
}
