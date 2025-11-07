package graphql

import (
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

type steps struct{}

func GetAllSteps() []stepbuilder.Step {
	st := steps{}

	return []stepbuilder.Step{
		st.prepareGraphQLRequest(),
		st.setGraphQLVariables(),
		st.setGraphQLArrayVariable(),
		st.setGraphQLHeaders(),
		st.sendGraphQLRequest(),
		st.validateGraphQLResponse(),
		st.validateGraphQLErrors(),
		st.validateGraphQLHasErrors(),
		st.validateGraphQLErrorMessage(),
		st.validateGraphQLDataValue(),
		st.validateGraphQLResponseDataShouldBe(),
		st.storeGraphQLData(),
		st.storeGraphQLArrayData(),
		st.storeGraphQLError(),
		st.storeGraphQLErrorMessage(),
	}
}

func GetDocs() []stepbuilder.Documentation {
	var docs []stepbuilder.Documentation
	for _, step := range GetAllSteps() {
		docs = append(docs, step.GetDocumentation())
	}
	return docs
}
