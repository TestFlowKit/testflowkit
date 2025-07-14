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
	Type              VarType
}

type StepCategory string

const (
	Form       StepCategory = "form"
	Visual     StepCategory = "visual"
	Keyboard   StepCategory = "keyboard"
	Navigation StepCategory = "navigation"
	Mouse      StepCategory = "mouse"
	RESTAPI    StepCategory = "restapi"
	Variable   StepCategory = "variable"
)

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
