package httpauth

import (
	"bytes"
	"io"
	"net/http"

	"testflowkit/pkg/formatter"
	"testflowkit/pkg/logger"
)

// DebugTransport is an http.RoundTripper that wraps a base transport and logs
// pretty-printed request and response bodies when debug mode is active.
//
// It is designed to sit outside AuthTransport in the transport chain:
//
//	DebugTransport → AuthTransport → BaseTransport
//
// Body streams are fully read and then restored so that downstream consumers
// (the actual HTTP round-trip and the application code reading the response)
// are unaffected.
type DebugTransport struct {
	// Base is the next RoundTripper in the chain. Defaults to
	// http.DefaultTransport when nil.
	Base http.RoundTripper

	// MaxBodySize is the body-size ceiling for pretty-printing. Bodies larger
	// than this value are replaced with a truncation notice in the log.
	// Zero means use formatter.DefaultMaxBodySize.
	MaxBodySize int64
}

// RoundTrip implements http.RoundTripper.
func (d *DebugTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	d.logRequest(req)

	resp, err := d.base().RoundTrip(req)
	if err != nil {
		return nil, err
	}

	d.logResponse(resp)
	return resp, nil
}

func (d *DebugTransport) logRequest(req *http.Request) {
	if req.Body == nil {
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		// Do not block the request; log what we have.
		logger.DebugFf("Request body read error: %v", err)
		return
	}

	// Restore the body so AuthTransport and the HTTP stack can still read it.
	req.Body = io.NopCloser(bytes.NewReader(body))
	req.GetBody = func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewReader(body)), nil
	}

	contentType := req.Header.Get("Content-Type")
	maxSize := d.effectiveMaxSize()
	logger.DebugFf("→ Request body (%s):\n%s", contentType, formatter.Format(contentType, body, maxSize))
}

func (d *DebugTransport) logResponse(resp *http.Response) {
	if resp.Body == nil {
		return
	}

	body, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		logger.DebugFf("Response body read error: %v", err)
		// Restore an empty body so callers don't receive a nil reader.
		resp.Body = io.NopCloser(bytes.NewReader(nil))
		return
	}

	// Restore the body so application code can still read the full response.
	resp.Body = io.NopCloser(bytes.NewReader(body))

	contentType := resp.Header.Get("Content-Type")
	maxSize := d.effectiveMaxSize()
	logger.DebugFf("← Response body (%s):\n%s", contentType, formatter.Format(contentType, body, maxSize))
}

func (d *DebugTransport) base() http.RoundTripper {
	if d.Base != nil {
		return d.Base
	}
	return http.DefaultTransport
}

func (d *DebugTransport) effectiveMaxSize() int64 {
	if d.MaxBodySize > 0 {
		return d.MaxBodySize
	}
	return formatter.DefaultMaxBodySize
}
