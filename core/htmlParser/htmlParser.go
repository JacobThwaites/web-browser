package htmlparser

import (
	"errors"
	"fmt"
	"strings"
)

type DomElement struct {
	tag        string
	text       string
	properties []string
	children   []DomElement
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
			// Doctype
			if char == "!" {
				if i+14 < len(httpBody) && strings.HasPrefix(string(httpBody[i-1:i+15]), "<!DOCTYPE html>") {
					tokens = append(tokens, Token{DoctypeToken, "html"})
					i += 14 // skip the rest of "<!DOCTYPE html>"
				} else {
					return []Token{}, errors.New("invalid doctype")
				}
				isTag = false
			} else if char == "/" { // End tag
				i++ // skip '/'
				tagType := ""
				for i < len(httpBody) && httpBody[i] != '>' {
					tagType += string(httpBody[i])
					i++
				}
				tokens = append(tokens, Token{EndTagToken, tagType})
				isTag = false
			} else { // Start tag
				tagType := ""

				for i < len(httpBody) && httpBody[i] != '>' && httpBody[i] != ' ' {
					tagType += string(httpBody[i])
					i++
				}

				for i < len(httpBody) && httpBody[i] != '>' {
					i++
				}
				tokens = append(tokens, Token{StartTagToken, tagType})
				isTag = false
			}
		} else if char == "<" {
			isTag = true
		} else {
			// text nodes ignored for now
			curr += char
		}
	}

	return tokens, nil
}


// func Tokenize(httpBody []byte) ([]Token, error) {
// 	tokens := []Token{}
// 	i := 0
// 	for i < len(httpBody) {
// 		if httpBody[i] == '<' {
// 			// Doctype or comment
// 			if i+1 < len(httpBody) && httpBody[i+1] == '!' {
// 				if i+15 <= len(httpBody) && strings.HasPrefix(string(httpBody[i:i+15]), "<!DOCTYPE html>") {
// 					tokens = append(tokens, Token{DoctypeToken, "html"})
// 					i += 15
// 					continue
// 				} else {
// 					return nil, errors.New("invalid doctype")
// 				}
// 			}

// 			// End tag
// 			if i+1 < len(httpBody) && httpBody[i+1] == '/' {
// 				j := i + 2
// 				for j < len(httpBody) && httpBody[j] != '>' {
// 					j++
// 				}
// 				tagType := string(httpBody[i+2:j])
// 				tokens = append(tokens, Token{EndTagToken, tagType})
// 				i = j + 1
// 				continue
// 			}

// 			// Start tag
// 			j := i + 1
// 			for j < len(httpBody) && httpBody[j] != '>' && httpBody[j] != ' ' {
// 				j++
// 			}
// 			tagType := string(httpBody[i+1:j])

// 			// Skip attributes (for now)
// 			for j < len(httpBody) && httpBody[j] != '>' {
// 				j++
// 			}
// 			tokens = append(tokens, Token{StartTagToken, tagType})
// 			i = j + 1
// 			continue
// 		}
// 		// Text node (ignore whitespace for now)
// 		i++
// 	}
// 	return tokens, nil
// }

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
