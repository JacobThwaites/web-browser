package htmlparser

func TreesEqual(tree1 DomElement, tree2 DomElement) bool {
	if tree1.token.Type != tree2.token.Type ||
		tree1.token.Data != tree2.token.Data ||
		!MapsEqual(tree1.token.Properties, tree2.token.Properties) {
		return false
	}

	if len(tree1.children) == 0 && len(tree2.children) == 0 {
		return true
	}

	if len(tree1.children) != len(tree2.children) {
		return false
	}

	for i := range tree1.children {
		if !TreesEqual(tree1.children[i], tree2.children[i]) {
			return false
		}
	}

	return true
}

func MapsEqual(a, b map[string]string) bool {
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

func DomTreeToString(tree DomElement) string {
    return domTreeToStringHelper(tree, "")
}

func domTreeToStringHelper(tree DomElement, indent string) string {
    if len(tree.children) == 0 {
        return indent + tree.token.Data
    }

    result := indent + tree.token.Data + "\n" + indent + "{\n"

    for _, child := range tree.children {
        result += domTreeToStringHelper(child, indent+"\t") + "\n"
    }

    result += indent + "}"

    return result
}