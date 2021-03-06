package file

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"testing"
)

// CopyFile and CopyFileWithPerms with invalid src
func TestCopyFileWithInvalidSrc(t *testing.T) {
	tempFolder, err := ioutil.TempDir("", "arksync-util-file-test")
	defer os.RemoveAll(tempFolder)
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := CopyFile("/invalid/file/path", path.Join(tempFolder, "dest"))
	if err == nil {
		t.Fatal("Should have fail to copy an invalid src file")
	}
	if bytes != 0 {
		t.Fatal("Should have written 0 bytes")
	}
	bytes, err = CopyFileSetPerms("/invalid/file/path", path.Join(tempFolder, "dest"), 0666)
	if err == nil {
		t.Fatal("Should have fail to copy an invalid src file (copy w/perms)")
	}
	if bytes != 0 {
		t.Fatal("Should have written 0 bytes (copy w/perms)")
	}

}

// CopyFile and CopyFileSetPerms with invalid dest
func TestCopyFileWithInvalidDest(t *testing.T) {
	tempFolder, err := ioutil.TempDir("", "arksync-util-file-test")
	defer os.RemoveAll(tempFolder)
	if err != nil {
		t.Fatal(err)
	}
	src := path.Join(tempFolder, "file")
	err = ioutil.WriteFile(src, []byte("content"), 0740)
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := CopyFile(src, path.Join(tempFolder, "/invalid/dest/path"))
	if err == nil {
		t.Fatal("Should have failed to copy an invalid src file")
	}
	if bytes != 0 {
		t.Fatal("Should have written 0 bytes")
	}
	bytes, err = CopyFileSetPerms(src, path.Join(tempFolder, "/invalid/dest/path"), 0666)
	if err == nil {
		t.Fatal("Should have failed to copy an invalid src file (copy w/perms)")
	}
	if bytes != 0 {
		t.Fatal("Should have written 0 bytes (copy w/perms)")
	}
}

// CopyFile and CopyFileSetPerms with same src and dest
func TestCopyFileWithSameSrcAndDest(t *testing.T) {
	tempFolder, err := ioutil.TempDir("", "arksync-util-file-test")
	defer os.RemoveAll(tempFolder)
	if err != nil {
		t.Fatal(err)
	}
	file := path.Join(tempFolder, "file")
	err = ioutil.WriteFile(file, []byte("content"), 0740)
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := CopyFile(file, file)
	if err != nil {
		t.Fatal(err)
	}
	if bytes != 0 {
		t.Fatal("Should have written 0 bytes as it is the same file.")
	}
	bytes, err = CopyFileSetPerms(file, file, 0666)
	if err != nil {
		t.Fatal("Copy with permissions failed, err:", err)
	}
	if bytes != 0 {
		t.Fatal("Should have written 0 bytes as it is the same file (copy w/perms).")
	}
}

// CopyFile with same src and dest but path is different and not clean
func TestCopyFileWithSameSrcAndDestWithPathNameDifferent(t *testing.T) {
	tempFolder, err := ioutil.TempDir("", "arksync-util-file-test")
	defer os.RemoveAll(tempFolder)
	if err != nil {
		t.Fatal(err)
	}
	testFolder := path.Join(tempFolder, "test")
	err = os.MkdirAll(testFolder, 0740)
	if err != nil {
		t.Fatal(err)
	}
	file := path.Join(testFolder, "file")
	sameFile := testFolder + "/../test/file"
	err = ioutil.WriteFile(file, []byte("content"), 0740)
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := CopyFile(file, sameFile)
	if err != nil {
		t.Fatal(err)
	}
	if bytes != 0 {
		t.Fatal("Should have written 0 bytes as it is the same file.")
	}
	bytes, err = CopyFileSetPerms(file, sameFile, 0666)
	if err != nil {
		t.Fatal("Error reported on copy w/perms, err:", err)
	}
	if bytes != 0 {
		t.Fatal("Should have written 0 bytes as it is the same file (copy w/perms).")
	}
}

func TestCopyFile(t *testing.T) {
	tempFolder, err := ioutil.TempDir("", "arksync-util-file-test")
	defer os.RemoveAll(tempFolder)
	if err != nil {
		t.Fatal(err)
	}
	src := path.Join(tempFolder, "src")
	dest := path.Join(tempFolder, "dest")
	ioutil.WriteFile(src, []byte("content"), 0777)
	ioutil.WriteFile(dest, []byte("destContent"), 0777)
	bytes, err := CopyFile(src, dest)
	if err != nil {
		t.Fatal(err)
	}
	if bytes != 7 {
		t.Fatalf("Should have written %d bytes but wrote %d", 7, bytes)
	}
	actual, err := ioutil.ReadFile(dest)
	if err != nil {
		t.Fatal(err)
	}
	if string(actual) != "content" {
		t.Fatalf("Dest content was '%s', expected '%s'", string(actual), "content")
	}
}

func TestCopyFileSetPerms(t *testing.T) {
	tempFolder, err := ioutil.TempDir("", "arksync-util-file-test2")
	defer os.RemoveAll(tempFolder)
	if err != nil {
		t.Fatal(err)
	}
	src := path.Join(tempFolder, "src")
	dest := path.Join(tempFolder, "dest")
	ioutil.WriteFile(src, []byte("content"), 0777)
	ioutil.WriteFile(dest, []byte("destContent"), 0777)
	bytes, err := CopyFileSetPerms(src, dest, 0775)
	if err != nil {
		t.Fatalf("Failed to copy file and set perms\n  src: %s\n  dest: %s\n  err: %s\n", src, dest, err)
	}
	if bytes != 7 {
		t.Fatalf("Should have written %d bytes but wrote %d", 7, bytes)
	}
	actual, err := ioutil.ReadFile(dest)
	if err != nil {
		t.Fatal(err)
	}
	if string(actual) != "content" {
		t.Fatalf("Dest content was '%s', expected '%s'", string(actual), "content")
	}
	info, _ := os.Stat(dest)
	mode := info.Mode()
	if mode != 0775 {
		t.Fatalf("Dest file permissions/mode was not 0755 as expected, found: %+v\n", mode)
	}
}

func TestWildcardMatches(t *testing.T) {
	match, _ := Matches("file.go", []string{"*"})
	if match != true {
		t.Errorf("failed to get a wildcard match, got %v", match)
	}
}

// A simple pattern match should return true.
func TestPatternMatches(t *testing.T) {
	match, _ := Matches("file.go", []string{"*.go"})
	if match != true {
		t.Errorf("failed to get a match, got %v", match)
	}
}

// An exclusion followed by an inclusion should return true.
func TestExclusionPatternMatchesPatternBefore(t *testing.T) {
	match, _ := Matches("file.go", []string{"!file.go", "*.go"})
	if match != true {
		t.Errorf("failed to get true match on exclusion pattern, got %v", match)
	}
}

// A folder pattern followed by an exception should return false.
func TestPatternMatchesFolderExclusions(t *testing.T) {
	match, _ := Matches("docs/README.md", []string{"docs", "!docs/README.md"})
	if match != false {
		t.Errorf("failed to get a false match on exclusion pattern, got %v", match)
	}
}

// A folder pattern followed by an exception should return false.
func TestPatternMatchesFolderWithSlashExclusions(t *testing.T) {
	match, _ := Matches("docs/README.md", []string{"docs/", "!docs/README.md"})
	if match != false {
		t.Errorf("failed to get a false match on exclusion pattern, got %v", match)
	}
}

// A folder pattern followed by an exception should return false.
func TestPatternMatchesFolderWildcardExclusions(t *testing.T) {
	match, _ := Matches("docs/README.md", []string{"docs/*", "!docs/README.md"})
	if match != false {
		t.Errorf("failed to get a false match on exclusion pattern, got %v", match)
	}
}

// A pattern followed by an exclusion should return false.
func TestExclusionPatternMatchesPatternAfter(t *testing.T) {
	match, _ := Matches("file.go", []string{"*.go", "!file.go"})
	if match != false {
		t.Errorf("failed to get false match on exclusion pattern, got %v", match)
	}
}

// A filename evaluating to . should return false.
func TestExclusionPatternMatchesWholeDirectory(t *testing.T) {
	match, _ := Matches(".", []string{"*.go"})
	if match != false {
		t.Errorf("failed to get false match on ., got %v", match)
	}
}

// A single ! pattern should return an error.
func TestSingleExclamationError(t *testing.T) {
	_, err := Matches("file.go", []string{"!"})
	if err == nil {
		t.Errorf("failed to get an error for a single exclamation point, got %v", err)
	}
}

// A string preceded with a ! should return true from Exclusion.
func TestExclusion(t *testing.T) {
	exclusion := exclusion("!")
	if !exclusion {
		t.Errorf("failed to get true for a single !, got %v", exclusion)
	}
}

// Matches with no patterns
func TestMatchesWithNoPatterns(t *testing.T) {
	matches, err := Matches("/any/path/there", []string{})
	if err != nil {
		t.Fatal(err)
	}
	if matches {
		t.Fatalf("Should not have match anything")
	}
}

// Matches with malformed patterns
func TestMatchesWithMalformedPatterns(t *testing.T) {
	matches, err := Matches("/any/path/there", []string{"["})
	if err == nil {
		t.Fatal("Should have failed because of a malformed syntax in the pattern")
	}
	if matches {
		t.Fatalf("Should not have match anything")
	}
}

// An empty string should return true from Empty.
func TestEmpty(t *testing.T) {
	empty := empty("")
	if !empty {
		t.Errorf("failed to get true for an empty string, got %v", empty)
	}
}

func TestCleanPatterns(t *testing.T) {
	cleaned, _, _, _ := CleanPatterns([]string{"docs", "config"})
	if len(cleaned) != 2 {
		t.Errorf("expected 2 element slice, got %v", len(cleaned))
	}
}

func TestCleanPatternsStripEmptyPatterns(t *testing.T) {
	cleaned, _, _, _ := CleanPatterns([]string{"docs", "config", ""})
	if len(cleaned) != 2 {
		t.Errorf("expected 2 element slice, got %v", len(cleaned))
	}
}

func TestCleanPatternsExceptionFlag(t *testing.T) {
	_, _, exceptions, _ := CleanPatterns([]string{"docs", "!docs/README.md"})
	if !exceptions {
		t.Errorf("expected exceptions to be true, got %v", exceptions)
	}
}

func TestCleanPatternsLeadingSpaceTrimmed(t *testing.T) {
	_, _, exceptions, _ := CleanPatterns([]string{"docs", "  !docs/README.md"})
	if !exceptions {
		t.Errorf("expected exceptions to be true, got %v", exceptions)
	}
}

func TestCleanPatternsTrailingSpaceTrimmed(t *testing.T) {
	_, _, exceptions, _ := CleanPatterns([]string{"docs", "!docs/README.md  "})
	if !exceptions {
		t.Errorf("expected exceptions to be true, got %v", exceptions)
	}
}

func TestCleanPatternsErrorSingleException(t *testing.T) {
	_, _, _, err := CleanPatterns([]string{"!"})
	if err == nil {
		t.Errorf("expected error on single exclamation point, got %v", err)
	}
}

func TestCleanPatternsFolderSplit(t *testing.T) {
	_, dirs, _, _ := CleanPatterns([]string{"docs/config/CONFIG.md"})
	if dirs[0][0] != "docs" {
		t.Errorf("expected first element in dirs slice to be docs, got %v", dirs[0][1])
	}
	if dirs[0][1] != "config" {
		t.Errorf("expected first element in dirs slice to be config, got %v", dirs[0][1])
	}
}

func TestCreateIfNotExistsFile(t *testing.T) {
	tempFolder, err := ioutil.TempDir("", "arksync-util-file-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempFolder)

	fileToCreate := filepath.Join(tempFolder, "file/to/create")

	if err := CreateIfNotExists(fileToCreate); err != nil {
		t.Fatal(err)
	}
	fileinfo, err := os.Stat(fileToCreate)
	if err != nil {
		t.Fatalf("Should have created a file, got %v", err)
	}

	if fileinfo.IsDir() {
		t.Fatalf("Should have been a file, seems it's not")
	}
}

