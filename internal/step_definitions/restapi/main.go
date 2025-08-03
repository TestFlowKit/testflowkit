package restapi

import (
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/internal/step_definitions/restapi/assertions"
)

type steps struct {
}

func GetAllSteps() []stepbuilder.Step {
	st := steps{}

	return []stepbuilder.Step{
		st.prepareRequest(),
		st.setHeaders(),
		st.setQueryParams(),
		st.setRequestBody(),
		st.setPathParams(),
		st.sendRequest(),
	}
}

func GetAssertionSteps() []stepbuilder.Step {
	return assertions.GetSteps()
}

func GetDocs() []stepbuilder.Documentation {
	var docs []stepbuilder.Documentation
	for _, step := range GetAllSteps() {
		docs = append(docs, step.GetDocumentation())
	}
	return docs
}
