package htmlparser

import (
	"errors"
	"fmt"
	"strings"
)

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
	Properties map[string]string
}

func NewToken(tokenType TokenType, data string, props ...map[string]string) Token {
	var properties map[string]string
    if len(props) > 0 {
        properties = props[0]
    }

    return Token{
        Type:       tokenType,
        Data:       data,
        Properties: properties,
    }
}

func NewDomElement(tokenType TokenType, data string, props ...map[string]string) DomElement {
	var properties map[string]string
    if len(props) > 0 {
        properties = props[0]
    }

    token := Token{
        Type:       tokenType,
        Data:       data,
        Properties: properties,
    }

	return DomElement{token, []DomElement{}}
}

func formatProperties(properties string) (map[string]string, error) {
	formatted := make(map[string]string)

	pairs := splitTagProperties(properties)

	for i := range pairs {
		property := pairs[i]

		if len(property) == 0 {
			continue
		}

		key := ""
		value := ""

		for j := 0; j < len(property); j++ {
			for j < len(property) {
				if property[j] == '=' {
					j++
					break
				}

				key += string(property[j])
				j++
			}

			if len(key) == 0 {
				return make(map[string]string), errors.New("invalid tag property")
			}

			if j >= len(property) {
				break
			}

			if property[j] != '"' || property[len(property) - 1] != '"' {
				return make(map[string]string), errors.New("invalid tag property key")
			}

			// ignore opening quotation mark
			j++

			for j < len(property) - 1 {
				value += string(property[j])
				j++
			}
		}

		formatted[key] = value
	}

	return formatted, nil
}

func splitTagProperties(tag string) []string {
	properties := []string{}

	property := ""
	isKey := true
	quoteCount := 0

	for _, char := range tag {
		if char == ' '  {
			if isKey {
				if len(property) == 0 {
					continue
				}

				properties = append(properties, property)
				property = ""
				continue
			}
		} else if char == '=' {
			isKey = false
		} else if char == '"' {
			quoteCount++

			if quoteCount == 2 {
				property += string(char)
				properties = append(properties, property)
				property = ""
				continue
			}
		}

		property += string(char)
	}

	return properties
}

func Tokenize(httpBody []byte) ([]Token, error) {
	tokens := []Token{}
	isTag := false

	for i := 0; i < len(httpBody); i++ {
		char := string(httpBody[i])
		if isTag {
			if char == "!" { // Doctype
				if i+14 < len(httpBody) && strings.HasPrefix(strings.ToLower(string(httpBody[i-1:i+15])), "<!doctype html>") {
					tokens = append(tokens, NewToken(DoctypeToken, "html"))
					i += 14 // skip the rest of "<!DOCTYPE html>"
				} else if i+3 < len(httpBody) && string(httpBody[i:i+3]) == "!--" { // Comment tag
					comment := ""
					i += 3
					for i+3 < len(httpBody) && string(httpBody[i:i+3]) != "-->" {
						comment += string(httpBody[i])
						i++
					}
					tokens = append(tokens, NewToken(CommentToken, comment))
					i += 3
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
				tokens = append(tokens, NewToken(EndTagToken, tagType))
				isTag = false
			} else { // Start tag
				tagType := ""

				for i < len(httpBody) && httpBody[i] != '>' && httpBody[i] != ' ' {
					tagType += string(httpBody[i])
					i++
				}

				// Get propeties
				properties := ""
				for i < len(httpBody) && httpBody[i] != '>' {
					properties += string(httpBody[i])
					i++
				}

				formattedProperties, err := formatProperties(properties)

				if err != nil {
					return []Token{}, err
				}

				tokens = append(tokens, NewToken(StartTagToken, tagType, formattedProperties))
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
				tokens = append(tokens, NewToken(TextToken, text))
			}
			i--
		}
	}

	return tokens, nil
}

type DomElement struct {
	token Token
	children []DomElement
}

func GenerateDomTree(tokens []Token) DomElement {
	head := DomElement{tokens[0], []DomElement{}}

	for i := 1; i < len(tokens); i++ {
		head, i = recursiveGenerate(tokens, head, i)
		fmt.Println(i)
	}

	return head
}

func recursiveGenerate(tokens []Token, curr DomElement, index int) (DomElement, int) {
    if index >= len(tokens) {
        return curr, index
    }

    for index < len(tokens) {
        token := tokens[index]

        // tags without children
        if token.Type == CommentToken || token.Type == TextToken {
            curr.children = append(curr.children, NewDomElement(token.Type, token.Data, token.Properties))
            index++
            continue
        }

        // closing tag for current level
        if token.Type == EndTagToken && token.Data == curr.token.Data {
            return curr, index + 1
        }

        // start a new element and recurse
        newElem := NewDomElement(token.Type, token.Data, token.Properties)
        child, newIndex := recursiveGenerate(tokens, newElem, index+1)
        curr.children = append(curr.children, child)
        index = newIndex
    }

    return curr, index
}


func ParseHTML(httpBody []byte) (DomElement, error) {
    s := strings.TrimSpace(string(httpBody))
    if !strings.HasPrefix(strings.ToLower(s), "<!doctype html>") {
        return DomElement{}, errors.New("invalid doctype")
    }

	tokens, err := Tokenize(httpBody)

	if err != nil {
		return DomElement{}, err
	}

	domTree := GenerateDomTree(tokens)

	return domTree, nil
}
