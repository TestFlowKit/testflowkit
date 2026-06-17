package configschema

import (
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"strings"
	"testflowkit/internal/config"
)

func loadDocIndex() (*docIndex, error) {
	idx := &docIndex{
		typeDocs:  make(map[string]string),
		fieldDocs: make(map[string]map[string]string),
	}

	fset := token.NewFileSet()

	for _, name := range config.SchemaSourceFileNames {
		if err := indexSourceFile(idx, fset, name); err != nil {
			return nil, err
		}
	}

	return idx, nil
}

func indexSourceFile(idx *docIndex, fset *token.FileSet, name string) error {
	content, err := config.SchemaSourceFiles.ReadFile(name)
	if err != nil {
		return fmt.Errorf("read embedded %s: %w", name, err)
	}

	file, err := parser.ParseFile(fset, name, content, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("parse %s: %w", name, err)
	}

	for _, decl := range file.Decls {
		genDecl, isGenDecl := decl.(*ast.GenDecl)
		if !isGenDecl || genDecl.Tok != token.TYPE {
			continue
		}
		indexTypeDecl(idx, genDecl)
	}

	return nil
}

func indexTypeDecl(idx *docIndex, genDecl *ast.GenDecl) {
	declDoc := ""
	if genDecl.Doc != nil {
		declDoc = strings.TrimSpace(genDecl.Doc.Text())
	}

	for _, spec := range genDecl.Specs {
		typeSpec, isTypeSpec := spec.(*ast.TypeSpec)
		if !isTypeSpec {
			continue
		}
		indexTypeSpec(idx, typeSpec, declDoc)
	}
}

func indexTypeSpec(idx *docIndex, typeSpec *ast.TypeSpec, declDoc string) {
	structName := typeSpec.Name.Name
	switch {
	case declDoc != "":
		idx.typeDocs[structName] = declDoc
	case typeSpec.Comment != nil:
		idx.typeDocs[structName] = strings.TrimSpace(typeSpec.Comment.Text())
	}

	structType, isStruct := typeSpec.Type.(*ast.StructType)
	if !isStruct {
		return
	}

	indexStructFields(idx, structName, structType)
}

func indexStructFields(idx *docIndex, structName string, structType *ast.StructType) {
	if idx.fieldDocs[structName] == nil {
		idx.fieldDocs[structName] = make(map[string]string)
	}

	for _, field := range structType.Fields.List {
		if len(field.Names) == 0 {
			continue
		}
		fieldName := field.Names[0].Name
		if !ast.IsExported(fieldName) {
			continue
		}

		var parts []string
		if field.Doc != nil {
			parts = append(parts, strings.TrimSpace(field.Doc.Text()))
		}
		if field.Comment != nil {
			parts = append(parts, strings.TrimSpace(field.Comment.Text()))
		}
		if len(parts) > 0 {
			idx.fieldDocs[structName][fieldName] = strings.Join(parts, "\n\n")
		}
	}
}

func splitDocComment(full string) (synopsis, details string) {
	full = strings.TrimSpace(full)
	if full == "" {
		return "", ""
	}
	synopsis = (&doc.Package{}).Synopsis(full)
	details = strings.TrimSpace(strings.TrimPrefix(full, synopsis))
	return synopsis, details
}
