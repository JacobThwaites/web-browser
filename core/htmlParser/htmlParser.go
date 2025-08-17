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

func generateDomTree(httpBody []byte) ([]DomElement, error) {
	fmt.Println("generate dom")
	fmt.Println(httpBody)

	tree := []DomElement{}
	return tree, nil
}

func ParseHTML(httpBody []byte) ([]DomElement, error) {

	docType := strings.ToLower(string(httpBody[:15]))

	if docType != "<!doctype html>" {
		return []DomElement{}, errors.New("invalid doctype")
	}

	domTree, err := generateDomTree(httpBody[15:])

	return domTree, err
}