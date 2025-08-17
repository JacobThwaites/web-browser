package htmlparser

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func openTestFile(filename string) ([]byte, error) {
	_, thisFile, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(thisFile)

	path := filepath.Join(currentDir, "testData", filename)

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func TestInvalidDocType(t *testing.T) {
	invalidDocTypeFile, err := openTestFile("invalidDoctype.html")

	if err != nil {
		t.Fatalf("failed to open test file: %v", err)
	}

	_, err = ParseHTML(invalidDocTypeFile)

	if err.Error() != "invalid doctype" {
		t.Errorf("Error Message = %v; want %v", err.Error(), "invalid doctype")
	}
}

// func TestGetDomain(t *testing.T) {
// 	// get directory of this test file
// 	_, filename, _, _ := runtime.Caller(0)
// 	currentDir := filepath.Dir(filename)
// 	path := filepath.Join(currentDir, "testData", "basic.html")

// 	f, err := os.Open(path)
// 	if err != nil {
// 		t.Fatalf("failed to open test file: %v", err)
// 	}

// 	defer f.Close()

// 	domTree, err := ParseHTML()

// 	if err != nil {
// 		t.Errorf("Error parsing HTML: %v", err)
// 	}

// 	fmt.Println(domTree)

// 	got := 1
//     if got != 1 {
//         t.Errorf("Abs(-1) = %d; want 1", got)
//     }
// }