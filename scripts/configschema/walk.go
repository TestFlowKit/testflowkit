package configschema

import "strings"

// CollectEmptyDescriptions returns YAML paths of schema nodes that lack a synopsis.
func CollectEmptyDescriptions(node *Node, prefix string) []string {
	if node == nil {
		return nil
	}

	var missing []string
	path := prefix
	if node.YAMLKey != "" {
		if prefix == "" {
			path = node.YAMLKey
		} else {
			path = prefix + "." + node.YAMLKey
		}
	}

	if node.Type == schemaTypeObject || node.Type == schemaTypeMap || node.Type == schemaTypeArray ||
		node.Type == schemaTypeString || node.Type == schemaTypeInteger || node.Type == schemaTypeBoolean ||
		node.Type == schemaTypeNumber {
		if node.Description == "" && node.YAMLKey != "" {
			missing = append(missing, path)
		}
	}

	for key, child := range node.Properties {
		childPath := path
		if childPath == "" {
			childPath = key
		} else {
			childPath = path + "." + key
		}
		missing = append(missing, CollectEmptyDescriptions(child, childPath)...)
	}
	if node.Value != nil {
		missing = append(missing, CollectEmptyDescriptions(node.Value, path+"<value>")...)
	}
	if node.Items != nil {
		missing = append(missing, CollectEmptyDescriptions(node.Items, path+"[]")...)
	}

	return missing
}

func findNodeByPath(node *Node, path string) *Node {
	if node == nil || path == "" {
		return node
	}
	parts := strings.Split(path, ".")
	current := node
	for _, part := range parts {
		if current == nil || current.Properties == nil {
			return nil
		}
		child, ok := current.Properties[part]
		if !ok {
			return nil
		}
		current = child
	}
	return current
}
