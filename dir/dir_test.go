package dir

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateIfNotExistsDirAndFindDir(t *testing.T) {
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

	folderForDvln := filepath.Join(tempFolder, ".dvln")
	if err := CreateIfNotExists(folderForDvln); err != nil {
		t.Fatal(err)
	}
	fileinfo, err = os.Stat(folderForDvln)
	if err != nil {
		t.Fatalf("Should have create a .dvln folder, got %v", err)
	}
	if !fileinfo.IsDir() {
		t.Fatalf("Should have been a dir, seems .dvln is not")
	}
	exists, err = Exists(folderForDvln)
	if err != nil {
		t.Fatal("Folder .dvln should have existed but instead an error was returned")
	}
	if !exists {
		t.Fatal("Folder .dvln should have existed but Exists() said it did not")
	}

	found, err := FindDirInOrAbove(folderToCreate, ".dvln_bogusname")
	if err != nil {
		t.Fatalf("Search for .dvln_bogusname should not have returned an error, got %v", err)
	}
	if found != "" {
		t.Fatal("Folder .dvln_bogusname should not have been found")
	}

	found, err = FindDirInOrAbove(folderToCreate, ".dvln")
	if err != nil {
		t.Fatalf("Search for .dvln dir should not have returned an error, got %v", err)
	}
	if found == "" {
		t.Fatal("Folder .dvln should not have been found but was not")
	}
	if found != tempFolder {
		t.Fatalf("Folder .dvln found in %s, should have been found in %s", found, tempFolder)
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
