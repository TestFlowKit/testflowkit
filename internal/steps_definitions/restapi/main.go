package restapi

import (
	"testflowkit/internal/steps_definitions/core/stepbuilder"
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
		st.checkResponseStatusCode(),
		st.responseBodyShouldContain(),
		st.responseBodyPathShouldExist(),
	}
}

func GetDocs() []stepbuilder.Documentation {
	var docs []stepbuilder.Documentation
	for _, step := range GetAllSteps() {
		docs = append(docs, step.GetDocumentation())
	}
	return docs
}
