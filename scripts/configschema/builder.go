package configschema

import "reflect"

func buildSchemaNode(
	t reflect.Type,
	typeName string,
	yamlKey string,
	required bool,
	docs *docIndex,
	validateTag string,
) *Node {
	t = dereferenceType(t)

	node := &Node{
		Name:     typeName,
		YAMLKey:  yamlKey,
		GoType:   t.String(),
		Required: required,
	}

	applyValidateTag(node, validateTag)
	applyHardcodedEnum(node, t)

	switch t.Kind() {
	case reflect.Struct:
		node.Type = schemaTypeObject
		applyStructDocumentation(node, t.Name(), docs)
		populateStructNode(node, t, docs)
	case reflect.Map:
		node.Type = schemaTypeMap
		node.KeyType = t.Key().Kind().String()
		node.Value = buildSchemaNode(t.Elem(), "", "", false, docs, "")
	case reflect.Slice, reflect.Array:
		node.Type = schemaTypeArray
		node.Items = buildSchemaNode(t.Elem(), "", "", false, docs, "")
	case reflect.Pointer:
		applyPointerSchema(node, t, typeName, yamlKey, docs, validateTag)
	case reflect.Bool:
		node.Type = schemaTypeBoolean
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		node.Type = schemaTypeInteger
	case reflect.Float32, reflect.Float64:
		node.Type = schemaTypeNumber
	case reflect.String:
		node.Type = schemaTypeString
	case reflect.Invalid, reflect.Uintptr, reflect.Complex64, reflect.Complex128,
		reflect.Chan, reflect.Func, reflect.Interface, reflect.UnsafePointer:
		node.Type = schemaTypeAny
	}

	return node
}

func applyValidateTag(node *Node, validateTag string) {
	if validateTag == "" {
		return
	}
	node.Constraints = validateTag
	if enum := parseOneOfEnum(validateTag); len(enum) > 0 {
		node.Enum = enum
	}
}

func applyHardcodedEnum(node *Node, t reflect.Type) {
	if t.Name() == "" || len(node.Enum) > 0 {
		return
	}
	if hardcoded := hardcodedEnum(t.Name()); len(hardcoded) > 0 {
		node.Enum = hardcoded
	}
}

func applyStructDocumentation(node *Node, structName string, docs *docIndex) {
	if structName == "" {
		return
	}
	typeDoc, ok := docs.typeDocs[structName]
	if !ok {
		return
	}
	synopsis, _ := splitDocComment(typeDoc)
	node.TypeDescription = synopsis
}

func applyPointerSchema(
	node *Node,
	t reflect.Type,
	typeName string,
	yamlKey string,
	docs *docIndex,
	validateTag string,
) {
	node.Required = false
	elem := t.Elem()
	node.GoType = t.String()

	if elem.Kind() != reflect.Struct {
		inner := buildSchemaNode(elem, typeName, yamlKey, false, docs, validateTag)
		node.Type = inner.Type
		node.Enum = inner.Enum
		node.Value = inner.Value
		node.Items = inner.Items
		return
	}

	node.Type = schemaTypeObject
	applyStructDocumentation(node, elem.Name(), docs)
	populateStructNode(node, elem, docs)
}

func populateStructNode(node *Node, t reflect.Type, docs *docIndex) {
	structName := t.Name()
	node.Properties = make(map[string]*Node)

	for i := range t.NumField() {
		field := t.Field(i)
		if !field.IsExported() {
			continue
		}

		yamlName, skip := parseYAMLTag(field.Tag.Get("yaml"))
		if skip {
			continue
		}

		validateTag := field.Tag.Get("validate")
		fieldRequired := isFieldRequired(field, validateTag)
		child := buildSchemaNode(field.Type, field.Name, yamlName, fieldRequired, docs, validateTag)

		if fieldDocs, structOk := docs.fieldDocs[structName]; structOk {
			if fieldDoc, fieldOk := fieldDocs[field.Name]; fieldOk {
				synopsis, details := splitDocComment(fieldDoc)
				child.Description = synopsis
				child.Details = details
			}
		}

		propKey := yamlName
		if propKey == "" {
			propKey = field.Name
		}
		node.Properties[propKey] = child
	}
}
