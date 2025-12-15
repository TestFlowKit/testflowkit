package scenario

import (
	"testflowkit/pkg/graphql"
)

func (bc *BackendContext) SetGraphQLRequest(request *graphql.Request) {
	bc.GraphQLRequest = request
}

func (bc *BackendContext) GetGraphQLRequest() *graphql.Request {
	return bc.GraphQLRequest
}

func (bc *BackendContext) GetGraphQLErrors() []graphql.Error {
	if bc.Response == nil {
		return nil
	}
	return bc.Response.GraphQLErrors
}

func (bc *BackendContext) HasGraphQLErrors() bool {
	return bc.Response != nil && len(bc.Response.GraphQLErrors) > 0
}
