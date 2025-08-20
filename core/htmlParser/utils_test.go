package htmlparser

import (
	"testing"
)

func TestTreesEqualTrue(t *testing.T) {
	tree1 := NewDomElement(DoctypeToken, "")
	htmlTag := NewDomElement(StartTagToken, "html")
	htmlTag.children = append(htmlTag.children, NewDomElement(TextToken, "some text"))
	tree1.children = append(tree1.children, htmlTag)
	tree1.children = append(tree1.children, NewDomElement(EndTagToken, "end"))

	tree2 := NewDomElement(DoctypeToken, "")
	htmlTag2 := NewDomElement(StartTagToken, "html")
	htmlTag2.children = append(htmlTag2.children, NewDomElement(TextToken, "some text"))
	tree2.children = append(tree2.children, htmlTag)
	tree2.children = append(tree2.children, NewDomElement(EndTagToken, "end"))
	equal := TreesEqual(tree1, tree2)

	if equal != true {
            t.Errorf("incorrect tree comparison: got %+v, want %+v", equal, true)
        }
}

func TestTreesEqualTextNotEquivalent(t *testing.T) {
	tree1 := NewDomElement(DoctypeToken, "")
	htmlTag := NewDomElement(StartTagToken, "html")
	htmlTag.children = append(htmlTag.children, NewDomElement(TextToken, "some text"))
	tree1.children = append(tree1.children, htmlTag)
	tree1.children = append(tree1.children, NewDomElement(EndTagToken, "end"))

	tree2 := NewDomElement(DoctypeToken, "")
	htmlTag2 := NewDomElement(StartTagToken, "html")
	htmlTag2.children = append(htmlTag2.children, NewDomElement(TextToken, "some textttttttt"))
	tree2.children = append(tree2.children, htmlTag2)
	tree2.children = append(tree2.children, NewDomElement(EndTagToken, "end"))
	equal := TreesEqual(tree1, tree2)

	if equal != false {
        t.Errorf("incorrect tree comparison, text not equivalent: got %+v, want %+v", equal, true)
    }
}

func TestTreesEqualTagNotEquivalent(t *testing.T) {
	tree1 := NewDomElement(DoctypeToken, "")
	htmlTag := NewDomElement(StartTagToken, "html")
	htmlTag.children = append(htmlTag.children, NewDomElement(TextToken, "some text"))
	tree1.children = append(tree1.children, htmlTag)
	tree1.children = append(tree1.children, NewDomElement(EndTagToken, "end"))

	tree2 := NewDomElement(DoctypeToken, "")
	htmlTag2 := NewDomElement(TextToken, "html")
	htmlTag2.children = append(htmlTag2.children, NewDomElement(TextToken, "some text"))
	tree2.children = append(tree2.children, htmlTag)
	tree2.children = append(tree2.children, NewDomElement(TextToken, "end"))
	equal := TreesEqual(tree1, tree2)

	if equal != false {
        t.Errorf("incorrect tree comparison, tags not equivalent: got %+v, want %+v", equal, true)
    }
}