package httpauth

import (
	"bytes"
	"io"
	"net/http"
	"time"

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

	start := time.Now()
	resp, err := d.base().RoundTrip(req)
	duration := time.Since(start)
	if err != nil {
		return nil, err
	}

	d.logResponse(resp, duration)
	return resp, nil
}

func (d *DebugTransport) logRequest(req *http.Request) {
	// Read and restore body if present
	var body []byte
	if req.Body != nil {
		b, err := io.ReadAll(req.Body)
		if err != nil {
			logger.DebugFf("Request body read error: %v", err)
		} else {
			body = b
		}

		req.Body = io.NopCloser(bytes.NewReader(body))
		req.GetBody = func() (io.ReadCloser, error) {
			return io.NopCloser(bytes.NewReader(body)), nil
		}
	}

	// Mask headers, URL, and body for logging
	maskedHeaders := logger.MaskHeaders(req.Header)
	maskedURL := logger.MaskURL(req.URL)
	contentType := req.Header.Get("Content-Type")
	maskedBody := logger.MaskBody(contentType, body)
	maxSize := d.effectiveMaxSize()

	logger.DebugFf("→ Request %s %s", req.Method, maskedURL)
	logger.DebugFf("Headers:\n%s", logger.HeadersToString(maskedHeaders))
	logger.DebugFf("Body (%s):\n%s", contentType, formatter.Format(contentType, maskedBody, maxSize))
}

func (d *DebugTransport) logResponse(resp *http.Response, duration time.Duration) {
	var body []byte
	if resp.Body != nil {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.DebugFf("Response body read error: %v", err)
			// Attempt to restore what we have and continue
			resp.Body = io.NopCloser(io.MultiReader(bytes.NewReader(b), resp.Body))
			return
		}
		_ = resp.Body.Close()
		resp.Body = io.NopCloser(bytes.NewReader(b))
		body = b
	}

	maskedHeaders := logger.MaskHeaders(resp.Header)
	contentType := resp.Header.Get("Content-Type")
	maskedBody := logger.MaskBody(contentType, body)

	logger.DebugFf("← Response %d %s (%s)", resp.StatusCode, resp.Status, duration)
	logger.DebugFf("Headers:\n%s", logger.HeadersToString(maskedHeaders))
	logger.DebugFf("Body (%s):\n%s", contentType, formatter.Format(contentType, maskedBody, d.effectiveMaxSize()))
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
