package htmlparser

import (
	"errors"
	"fmt"
	"strings"
)

type DomElement struct {
	tag string
	text string
	properties []string
	children []DomElement
}

type DomTree struct {

}

type TokenType int

const (
    ErrorToken TokenType = iota // special: error or end-of-file
    DoctypeToken
    StartTagToken
    EndTagToken
    SelfClosingTagToken
    TextToken
    CommentToken
)

func (tt TokenType) String() string {
    switch tt {
    case ErrorToken:
        return "ErrorToken"
    case DoctypeToken:
        return "DoctypeToken"
    case StartTagToken:
        return "StartTagToken"
    case EndTagToken:
        return "EndTagToken"
    case SelfClosingTagToken:
        return "SelfClosingTagToken"
    case TextToken:
        return "TextToken"
    case CommentToken:
        return "CommentToken"
    default:
        return "UnknownToken"
    }
}

type Token struct {
    Type TokenType
    Data string
}

func Tokenize(httpBody []byte) ([]Token, error) {
	tokens := []Token{}

	curr := ""
	isTag := false

	for i := 0; i < len(httpBody); i++ {
		char := string(httpBody[i])
		if isTag {
			curr += char
			if char == "!" { // doctype or comment
				if string(httpBody[i+1]) == "D" {
					if string(httpBody[i-1:i+14]) != "<!DOCTYPE html>" {
						return []Token{}, errors.New("invalid doctype")
					}

					tokens = append(tokens, Token{DoctypeToken, "html"})
					i += 14
					isTag = false
				} else { // comment

				}
			} else if char == "/" { // endtag
				tagType := ""
				for string(httpBody[i]) != ">" {
					tagType += string(httpBody[i])
					i++
				}

				tokens = append(tokens, Token{EndTagToken, tagType})
				isTag = false
			} else { // starttag
				tagType := ""
				for string(httpBody[i]) != ">" {
					tagType += string(httpBody[i])
					i++
				}

				tokens = append(tokens, Token{StartTagToken, tagType})
			}
		} else if char == "<" {
			isTag = true
		} else { // is text

		}
	}

	return tokens, nil
}


func ParseHTML(httpBody []byte) (DomElement, error) {
	docType := strings.ToLower(string(httpBody[:15]))

	if docType != "<!doctype html>" {
		return DomElement{}, errors.New("invalid doctype")
	}

	fmt.Println("calling func")

	// domTree, err := generateDomTree(httpBody[16:], DomElement{})
	domTree := DomElement{}

	return domTree, nil
}