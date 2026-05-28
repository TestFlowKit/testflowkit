package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// parseConfigComments parses all non-test Go files in dir and returns
// map[structName][goFieldName] = leading doc comment text.
func parseConfigComments(dir string) (map[string]map[string]string, error) {
	fset := token.NewFileSet()
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	result := make(map[string]map[string]string)
	for _, entry := range entries {
		if !isConfigSourceFile(entry) {
			continue
		}

		path := filepath.Join(dir, entry.Name())
		if errP := parseCommentFile(fset, path, result); errP != nil {
			return nil, err
		}
	}

	return result, nil
}

func isConfigSourceFile(entry os.DirEntry) bool {
	return !entry.IsDir() && strings.HasSuffix(entry.Name(), ".go") && !strings.HasSuffix(entry.Name(), "_test.go")
}

func parseCommentFile(fset *token.FileSet, path string, result map[string]map[string]string) error {
	file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	for _, decl := range file.Decls {
		collectStructComments(decl, result)
	}

	return nil
}

func collectStructComments(decl ast.Decl, result map[string]map[string]string) {
	genDecl, ok := decl.(*ast.GenDecl)
	if !ok {
		return
	}

	for _, spec := range genDecl.Specs {
		typeSpec, isSpec := spec.(*ast.TypeSpec)
		if !isSpec {
			continue
		}

		structType, isStruct := typeSpec.Type.(*ast.StructType)
		if !isStruct {
			continue
		}

		structName := typeSpec.Name.Name
		if result[structName] == nil {
			result[structName] = make(map[string]string)
		}

		addFieldComments(result[structName], structType)
	}
}

func addFieldComments(dst map[string]string, structType *ast.StructType) {
	for _, field := range structType.Fields.List {
		if len(field.Names) == 0 {
			continue
		}

		comment := fieldComment(field)
		for _, name := range field.Names {
			dst[name.Name] = comment
		}
	}
}

func fieldComment(field *ast.Field) string {
	if field.Doc != nil {
		return cleanComment(field.Doc.Text())
	}
	if field.Comment != nil {
		return cleanComment(field.Comment.Text())
	}
	return ""
}

// cleanComment trims whitespace and removes separator lines.
func cleanComment(raw string) string {
	lines := strings.Split(strings.TrimSpace(raw), "\n")
	var parts []string
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l == "" || strings.HasPrefix(l, "---") || strings.HasPrefix(l, "===") {
			continue
		}
		parts = append(parts, l)
	}
	return strings.Join(parts, " ")
}
