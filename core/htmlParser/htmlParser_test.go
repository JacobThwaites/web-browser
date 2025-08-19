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

func mapsEqual(a, b map[string]string) bool {
    if len(a) != len(b) {
        return false
    }
    for k, v := range a {
        if bv, ok := b[k]; !ok || bv != v {
            return false
        }
    }
    return true
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
        {Type: StartTagToken, Data: "html", Properties: map[string]string{"lang": "en"}},
		{Type: CommentToken, Data: " asdf->"},
        {Type: StartTagToken, Data: "head", Properties: map[string]string{"noValue": "", "foo": "bar"}},
        {Type: EndTagToken, Data: "head"},
        {Type: StartTagToken, Data: "body"},
		{Type: TextToken, Data: "Some text"},
        {Type: EndTagToken, Data: "body"},
        {Type: EndTagToken, Data: "html"},
    }

	if len(tokens) != len(expected) {
        t.Fatalf("token count mismatch: got %d, want %d", len(tokens), len(expected))
    }

    for i, tok := range tokens {
        if tok.Type != expected[i].Type || tok.Data != expected[i].Data || !mapsEqual(tok.Properties, expected[i].Properties) {
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