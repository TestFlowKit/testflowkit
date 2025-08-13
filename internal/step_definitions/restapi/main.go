package restapi

import (
	"slices"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/internal/step_definitions/restapi/assertions"
)

type steps struct {
}

func GetAllSteps() []stepbuilder.Step {
	st := steps{}

	return slices.Concat([]stepbuilder.Step{
		st.prepareRequest(),
		st.setHeaders(),
		st.setQueryParams(),
		st.setRequestBody(),
		st.setJSONRequestBody(),
		st.setPathParams(),
		st.debugRequest(),
		st.sendRequest(),
	}, GetAssertionSteps())
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
