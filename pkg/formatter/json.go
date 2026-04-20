package formatter

import (
	"bytes"
	"encoding/json"
)

// FormatJSON pretty-prints raw JSON bytes with 2-space indentation.
// If raw is not valid JSON the original bytes are returned together with the
// parse error so callers can fall back gracefully.
func FormatJSON(raw []byte) ([]byte, error) {
	var buf bytes.Buffer
	if err := json.Indent(&buf, raw, "", "  "); err != nil {
		return raw, err
	}
	return buf.Bytes(), nil
}
