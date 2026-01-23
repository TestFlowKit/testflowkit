package scenario

import (
	"testflowkit/pkg/graphql"
)

func (bc *BackendContext) SetGraphQLRequest(request *graphql.Request) {
	bc.GraphQL.Request = request
}

func (bc *BackendContext) GetGraphQLRequest() *graphql.Request {
	return bc.GraphQL.Request
}

func (bc *BackendContext) SetGraphQLEndpoint(endpoint string) {
	bc.GraphQL.Endpoint = endpoint
}

func (bc *BackendContext) GetGraphQLEndpoint() string {
	return bc.GraphQL.Endpoint
}

func (bc *BackendContext) SetGraphQLHeader(key, value string) {
	if bc.GraphQL.Headers == nil {
		bc.GraphQL.Headers = make(map[string]string)
	}
	bc.GraphQL.Headers[key] = value
}

func (bc *BackendContext) GetGraphQLHeaders() map[string]string {
	if bc.GraphQL.Headers == nil {
		return make(map[string]string)
	}
	return bc.GraphQL.Headers
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
