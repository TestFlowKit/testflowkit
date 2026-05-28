package main

import (
	"reflect"
	"strings"
)

// reflectStructFields builds a []Field by walking the struct at t using yaml/validate tags
// and AST-extracted doc comments.
func reflectStructFields(t reflect.Type, comments map[string]map[string]string) []Field {
	t = deref(t)
	if t.Kind() != reflect.Struct {
		return nil
	}

	structComments := comments[t.Name()]
	var fields []Field

	for i := range t.NumField() {
		sf := t.Field(i)
		field, ok := buildField(sf, structComments, comments)
		if !ok {
			continue
		}
		fields = append(fields, field)
	}
	return fields
}

func buildField(
	sf reflect.StructField,
	structComments map[string]string,
	comments map[string]map[string]string,
) (Field, bool) {
	if !sf.IsExported() {
		return Field{}, false
	}

	yamlKey, skip := parseYAMLTag(sf.Tag.Get("yaml"))
	if skip {
		return Field{}, false
	}

	ft := deref(sf.Type)
	c := parseValidateTag(sf.Tag.Get("validate"))
	field := Field{
		Key:           yamlKey,
		Type:          goTypeToSchemaType(ft),
		Required:      c.required,
		RequiredIf:    c.requiredIf,
		Min:           c.min,
		Max:           c.max,
		Enum:          c.enum,
		Interpolation: ft.Kind() == reflect.String,
		Description:   fieldDescription(sf.Name, structComments),
	}

	attachNestedFields(&field, ft, comments)
	return field, true
}

func fieldDescription(name string, structComments map[string]string) string {
	if structComments == nil {
		return ""
	}

	return structComments[name]
}

func attachNestedFields(field *Field, ft reflect.Type, comments map[string]map[string]string) {
	if ft.Kind() == reflect.Struct {
		if nested := reflectStructFields(ft, comments); len(nested) > 0 {
			field.Type = schemaTypeObject
			field.Fields = nested
		}
		return
	}

	if ft.Kind() != reflect.Map {
		return
	}

	elem := deref(ft.Elem())
	if elem.Kind() != reflect.Struct {
		return
	}

	if nested := reflectStructFields(elem, comments); len(nested) > 0 {
		field.Fields = nested
	}
}

func parseYAMLTag(tag string) (key string, skip bool) {
	if tag == "" {
		return "", true
	}
	key = strings.Split(tag, ",")[0]
	if key == "" || key == "-" {
		return "", true
	}
	return key, false
}

func fieldMapByKey(fields []Field) map[string]Field {
	m := make(map[string]Field, len(fields))
	for _, f := range fields {
		m[f.Key] = f
	}
	return m
}
