package main

import (
	"fmt"
	"reflect"
)

const (
	schemaTypeString = "string"
	schemaTypeBool   = "bool"
	schemaTypeInt    = "int"
	schemaTypeInt64  = "int64"
	schemaTypeUint   = "uint"
	schemaTypeFloat  = "float"
	schemaTypeAny    = "any"
	schemaTypeObject = "object"
)

func goTypeToSchemaType(t reflect.Type) string {
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	switch t.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.String:
		return schemaTypeString
	case reflect.Bool:
		return schemaTypeBool
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		return schemaTypeInt
	case reflect.Int64:
		return schemaTypeInt64
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return schemaTypeUint
	case reflect.Float32, reflect.Float64:
		return schemaTypeFloat
	case reflect.Complex64, reflect.Complex128:
		return "complex"
	case reflect.Slice:
		return "[]" + goTypeToSchemaType(t.Elem())
	case reflect.Array:
		return "[]" + goTypeToSchemaType(t.Elem())
	case reflect.Map:
		key := goTypeToSchemaType(t.Key())
		elem := deref(t.Elem())
		if elem.Kind() == reflect.Struct && elem.Name() != "" {
			return fmt.Sprintf("map<%s, %s>", schemaTypeString, elem.Name())
		}
		return fmt.Sprintf("map<%s, %s>", key, goTypeToSchemaType(elem))
	case reflect.Struct:
		if t.Name() != "" {
			return t.Name()
		}
		return schemaTypeObject
	case reflect.Interface:
		return schemaTypeAny
	case reflect.Pointer:
		return goTypeToSchemaType(t.Elem())
	case reflect.Chan, reflect.Func, reflect.UnsafePointer:
		return t.Kind().String()
	}

	return t.Kind().String()
}

func deref(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	return t
}
