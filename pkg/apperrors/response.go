package apperrors

import "errors"

// ErrNoResponseAvailable is returned when a step expects a response but none has been received yet.
var ErrNoResponseAvailable = errors.New("no response available - send a request first")

// ErrNoGraphQLResponse is returned when a GraphQL-specific step expects a response but none is present.
var ErrNoGraphQLResponse = errors.New("no GraphQL response available - send a request first")

// ErrNoRESTResponse is returned when a REST-specific step expects a response but none is present.
var ErrNoRESTResponse = errors.New("no REST API response available")

// ErrNoGraphQLErrors is returned when expecting GraphQL errors but the response has none.
var ErrNoGraphQLErrors = errors.New("expected GraphQL errors but found none")

// ErrNoRequestPrepared is returned when a send step is executed without a prior prepare step.
var ErrNoRequestPrepared = errors.New("no request has been prepared - use 'I prepare a request' step first")
