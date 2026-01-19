package stepbuilder

import (
	"slices"
	"strings"
)

type Documentation struct {
	Sentence    string
	Description string
	Variables   []DocVariable
	Example     string
	Categories  []StepCategory
}

type DocVariable struct {
	Name, Description string
	Type              VarType
}

type StepCategory string

var Backend = []StepCategory{RESTAPI, GraphQL}

const (
	Form       StepCategory = "form"
	Visual     StepCategory = "visual"
	Keyboard   StepCategory = "keyboard"
	Navigation StepCategory = "navigation"
	Mouse      StepCategory = "mouse"
	RESTAPI    StepCategory = "restapi"
	GraphQL    StepCategory = "graphql"
	Variable   StepCategory = "variable"
	Assertions StepCategory = "assertions"
)

func mergeCategories(doc DocParams) []StepCategory {
	unique := make([]StepCategory, 0, len(doc.Categories))
	for _, c := range doc.Categories {
		if c == "" || slices.Contains(unique, c) {
			continue
		}
		unique = append(unique, c)
	}

	return unique
}

type VarType string

const (
	VarTypeString VarType = "string"
	VarTypeInt    VarType = "int"
	VarTypeFloat  VarType = "float"
	VarTypeBool   VarType = "bool"
	VarTypeTable  VarType = "table"
)

func VarTypeEnum(values ...string) VarType {
	return VarType(strings.Join(values, ", "))
}
