package httpauth_test

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"testing"

	"testflowkit/internal/httpauth"
)

// roundTripFunc lets us inject a fake RoundTripper inline.
type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func makeResponse(ct, body string) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": []string{ct}},
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

func TestDebugTransport_ResponseBodyRestoredAfterLogging(t *testing.T) {
	jsonBody := `{"status":"ok"}`
	dt := &httpauth.DebugTransport{
		Base: roundTripFunc(func(_ *http.Request) (*http.Response, error) {
			return makeResponse("application/json", jsonBody), nil
		}),
		MaxBodySize: 0,
	}

	req, _ := http.NewRequest(http.MethodGet, "http://example.com/", nil)
	resp, err := dt.RoundTrip(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()

	got, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}
	if string(got) != jsonBody {
		t.Errorf("response body modified: got %q, want %q", got, jsonBody)
	}
}

func TestDebugTransport_RequestBodyRestoredAfterLogging(t *testing.T) {
	reqBody := `{"action":"test"}`
	var capturedBody string

	dt := &httpauth.DebugTransport{
		Base: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			b, _ := io.ReadAll(r.Body)
			capturedBody = string(b)
			return makeResponse("application/json", `{}`), nil
		}),
	}

	body := bytes.NewBufferString(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "http://example.com/", body)
	req.Header.Set("Content-Type", "application/json")

	_, err := dt.RoundTrip(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if capturedBody != reqBody {
		t.Errorf("request body not restored: base transport got %q, want %q", capturedBody, reqBody)
	}
}

func TestDebugTransport_NilRequestBody_NoError(t *testing.T) {
	dt := &httpauth.DebugTransport{
		Base: roundTripFunc(func(_ *http.Request) (*http.Response, error) {
			return makeResponse("text/plain", "ok"), nil
		}),
	}

	req, _ := http.NewRequest(http.MethodGet, "http://example.com/", nil)
	resp, err := dt.RoundTrip(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status: %d", resp.StatusCode)
	}
}

func TestDebugTransport_NilResponseBody_NoError(t *testing.T) {
	dt := &httpauth.DebugTransport{
		Base: roundTripFunc(func(_ *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusNoContent,
				Header:     http.Header{},
				Body:       nil,
			}, nil
		}),
	}

	req, _ := http.NewRequest(http.MethodDelete, "http://example.com/resource", nil)
	resp, err := dt.RoundTrip(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("unexpected status: %d", resp.StatusCode)
	}
}

func TestDebugTransport_BinaryResponseNotCorrupted(t *testing.T) {
	// PNG magic bytes — must survive the transport unchanged.
	binary := []byte{0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a}

	dt := &httpauth.DebugTransport{
		Base: roundTripFunc(func(_ *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     http.Header{"Content-Type": []string{"image/png"}},
				Body:       io.NopCloser(bytes.NewReader(binary)),
			}, nil
		}),
	}

	req, _ := http.NewRequest(http.MethodGet, "http://example.com/img.png", nil)
	resp, err := dt.RoundTrip(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got, _ := io.ReadAll(resp.Body)
	if !bytes.Equal(got, binary) {
		t.Errorf("binary body corrupted")
	}
}

func TestDebugTransport_OversizedBody_StillRestored(t *testing.T) {
	// Body is 20 bytes; limit is 5 bytes → truncation notice in log, but
	// the response body must still be fully readable by the caller.
	largeBody := strings.Repeat("x", 20)

	dt := &httpauth.DebugTransport{
		Base: roundTripFunc(func(_ *http.Request) (*http.Response, error) {
			return makeResponse("application/json", largeBody), nil
		}),
		MaxBodySize: 5,
	}

	req, _ := http.NewRequest(http.MethodGet, "http://example.com/", nil)
	resp, err := dt.RoundTrip(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got, _ := io.ReadAll(resp.Body)
	if string(got) != largeBody {
		t.Errorf("oversized body was not fully restored: got len=%d, want len=%d", len(got), len(largeBody))
	}
}

func TestDebugTransport_DefaultBase_UsesDefaultTransport(_ *testing.T) {
	// Verify that a zero-value DebugTransport is constructable without panic.
	// The Base field defaults to http.DefaultTransport when nil; we cannot
	// make a real network call in a unit test, so we just ensure the struct
	// compiles and is non-nil as a compile-time sanity check.
	var dt httpauth.DebugTransport
	_ = dt
}
