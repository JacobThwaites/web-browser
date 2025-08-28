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
        {Type: StartTagToken, Data: "head", Properties: map[string]string{"noValue": "", "foo": "bar asdf"}},
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
        if tok.Type != expected[i].Type || tok.Data != expected[i].Data || !MapsEqual(tok.Properties, expected[i].Properties) {
            t.Errorf("token %d mismatch: got %+v, want %+v", i, tok, expected[i])
        }
    }
}

func TestGenerateDomTree(t *testing.T) {
	tokens := []Token{
		NewToken(DoctypeToken, "html"),
		NewToken(StartTagToken, "html"),
		NewToken(StartTagToken, "head"),
		NewToken(TextToken, "Some text"),
		NewToken(EndTagToken, "head"),
		NewToken(EndTagToken, "html"),
	}

	domTree := GenerateDomTree(tokens)

	expected :=  NewDomElement(DoctypeToken, "html")
	htmlTag := NewDomElement(StartTagToken, "html")
	headTag := NewDomElement(StartTagToken, "head")
	headTag.children = append(headTag.children, NewDomElement(TextToken, "Some text"))
	htmlTag.children = append(htmlTag.children, headTag)
	expected.children = append(expected.children, htmlTag)


	if !TreesEqual(domTree, expected) {
        t.Errorf("incorrect tree comparison, text not equivalent: got %+v, want %+v", DomTreeToString(domTree), DomTreeToString(expected))
    }
}
