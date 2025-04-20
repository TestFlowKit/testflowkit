package shared

import "strings"

type StepDocumentation struct {
	Sentence    string
	Description string
	Variables   []StepVariable
	Example     string
	Category    TestCategory
}

type StepVariable struct {
	Name, Description string
	Type              DocVarType
}

type TestCategory string

const (
	Form       TestCategory = "form"
	Visual     TestCategory = "visual"
	Keyboard   TestCategory = "keyboard"
	Navigation TestCategory = "navigation"
	Mouse      TestCategory = "mouse"
)

type DocVarType string

const (
	DocVarTypeString DocVarType = "string"
	DocVarTypeInt    DocVarType = "int"
	DocVarTypeFloat  DocVarType = "float"
	DocVarTypeBool   DocVarType = "bool"
	DocVarTypeTable  DocVarType = "table"
)

func DocVarTypeEnum(values ...string) DocVarType {
	return DocVarType(strings.Join(values, ", "))
}
