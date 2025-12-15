package protocol

import (
	"context"
)

type APIProtocol interface {
	PrepareRequest(ctx context.Context, name string) (context.Context, error)

	SendRequest(ctx context.Context) (context.Context, error)

	GetResponseBody(ctx context.Context) ([]byte, error)

	GetStatusCode(ctx context.Context) (int, error)

	HasErrors(ctx context.Context) bool

	GetProtocolName() string
}

type APIProtocolName string

const (
	ProtocolGraphQL APIProtocolName = "GraphQL"
	ProtocolRESTAPI APIProtocolName = "REST"
)
