package formatter

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
)

// FormatXML pretty-prints raw XML bytes with 2-space indentation using a
// token-based approach to preserve element content and attributes faithfully.
// If raw is not valid XML the original bytes are returned together with the
// parse error so callers can fall back gracefully.
func FormatXML(raw []byte) ([]byte, error) {
	var buf bytes.Buffer
	dec := xml.NewDecoder(bytes.NewReader(raw))
	enc := xml.NewEncoder(&buf)
	enc.Indent("", "  ")

	for {
		tok, err := dec.Token()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return raw, fmt.Errorf("xml decode: %w", err)
		}
		if err = enc.EncodeToken(tok); err != nil {
			return raw, fmt.Errorf("xml encode: %w", err)
		}
	}

	if err := enc.Flush(); err != nil {
		return raw, fmt.Errorf("xml flush: %w", err)
	}

	return buf.Bytes(), nil
}
