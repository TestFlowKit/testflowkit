package stepbuilder

import "strings"

type Documentation struct {
	Sentence    string
	Description string
	Variables   []DocVariable
	Example     string
	Category    StepCategory
}

type DocVariable struct {
	Name, Description string
	Type              DocVarType
}

type StepCategory string

const (
	Form       StepCategory = "form"
	Visual     StepCategory = "visual"
	Keyboard   StepCategory = "keyboard"
	Navigation StepCategory = "navigation"
	Mouse      StepCategory = "mouse"
)

type DocVarType string

const (
	VarTypeString DocVarType = "string"
	VarTypeInt    DocVarType = "int"
	VarTypeFloat  DocVarType = "float"
	VarTypeBool   DocVarType = "bool"
	VarTypeTable  DocVarType = "table"
)

func DocVarTypeEnum(values ...string) DocVarType {
	return DocVarType(strings.Join(values, ", "))
}
