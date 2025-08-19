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
	isTag := false

	for i := 0; i < len(httpBody); i++ {
		char := string(httpBody[i])
		if isTag {
			// Doctype
			if char == "!" {
				if i+14 < len(httpBody) && strings.HasPrefix(string(httpBody[i-1:i+15]), "<!DOCTYPE html>") {
					tokens = append(tokens, Token{DoctypeToken, "html"})
					i += 14 // skip the rest of "<!DOCTYPE html>"
				} else if i+3 < len(httpBody) && string(httpBody[i:i+3]) == "!--" { // Comment tag
					comment := ""
					i += 3
					for i+3 < len(httpBody) && string(httpBody[i:i+3]) != "-->" {
						comment += string(httpBody[i])
						i++
					}
					tokens = append(tokens, Token{CommentToken, comment})
					i += 3
				} else {
					fmt.Println(string(httpBody[i:i+2]))
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
			if isTag {
				return []Token{}, errors.New("invalid tag; double << ")
			}
			isTag = true
		} else {
			text := ""
			for i < len(httpBody) && string(httpBody[i]) != "<" {
				text += string(httpBody[i])
				i++
			}

			text = strings.TrimSpace(text)
			if text != "" {
				tokens = append(tokens, Token{TextToken, text})
			}
			i--
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
