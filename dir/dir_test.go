package dir

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateIfNotExistsDir(t *testing.T) {
	tempFolder, err := ioutil.TempDir("", "dvln-util-dir-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempFolder)

	folderToCreate := filepath.Join(tempFolder, "tocreate")

	if err := CreateIfNotExists(folderToCreate); err != nil {
		t.Fatal(err)
	}
	fileinfo, err := os.Stat(folderToCreate)
	if err != nil {
		t.Fatalf("Should have create a folder, got %v", err)
	}
	if !fileinfo.IsDir() {
		t.Fatalf("Should have been a dir, seems it's not")
	}

	exists, err := Exists(folderToCreate)
	if err != nil {
		t.Fatal("Folder should have existed but instead an error was returned")
	}
	if !exists {
		t.Fatal("Folder should have existed but Exists() said it did not")
	}
}

func TestExistsIfNotExists(t *testing.T) {
	exists, err := Exists(".boguselement")
	if err != nil {
		t.Fatal("File should not have existed but that should be normal and not an error")
	}
	if exists {
		t.Fatal("File should not have existed but Exists() found it")
	}
}
