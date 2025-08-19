package htmlparser

import (
	"fmt"
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

func TestTokenizer(t *testing.T) {
	file, err := openTestFile("simple.html")

	if err != nil {
		t.Fatalf("failed to open test file: %v", err)
	}

	tokens, err := Tokenize(file)

	if err != nil {
		t.Errorf("Error Message = %v; want %v", err.Error(), "invalid doctype")
	}

	expected := []Token{
        {Type: DoctypeToken, Data: "html"},
        {Type: StartTagToken, Data: "html"},
		{Type: CommentToken, Data: " asdf->"},
        {Type: StartTagToken, Data: "head"},
        {Type: EndTagToken, Data: "head"},
        {Type: StartTagToken, Data: "body"},
        {Type: EndTagToken, Data: "body"},
        {Type: EndTagToken, Data: "html"},
    }

	if len(tokens) != len(expected) {
        t.Fatalf("token count mismatch: got %d, want %d", len(tokens), len(expected))
    }

    for i, tok := range tokens {
        if tok.Type != expected[i].Type || tok.Data != expected[i].Data {
            t.Errorf("token %d mismatch: got %+v, want %+v", i, tok, expected[i])
        }
    }
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

func TestGetDomain(t *testing.T) {
	testFile, err := openTestFile("basic.html")

	if err != nil {
		t.Fatalf("failed to open test file: %v", err)
	}


	domTree, err := ParseHTML(testFile)

	if err != nil {
		fmt.Println(domTree)
		t.Errorf("Error parsing HTML: %v", err)
	}

}