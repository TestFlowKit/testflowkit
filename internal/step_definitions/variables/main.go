package variables

import (
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

type steps struct{}

func GetAllSteps() []stepbuilder.Step {
	st := steps{}
	return []stepbuilder.Step{
		st.storeJSONPathIntoVariable(),
		st.storeElementContentIntoVariable(),
		st.storeCustomVariable(),
		st.variableShouldContains(),
		st.storeValueIntoGlobalVariable(),
		st.storeJSONPathIntoGlobalVariable(),
	}
}

func GetDocs() []stepbuilder.Documentation {
	var docs []stepbuilder.Documentation
	for _, step := range GetAllSteps() {
		docs = append(docs, step.GetDocumentation())
	}
	return docs
}
