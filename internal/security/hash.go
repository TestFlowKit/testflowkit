package security

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"

	"testflowkit/internal/config"
)

// SchemeHash returns a deterministic, hex-encoded SHA-256 digest of a
// fully-resolved SecurityScheme.
//
// All {{ env.* }} references are already substituted before the config struct
// is populated (see internal/config/loader.go), so the resulting hash changes
// automatically whenever environment values change – this is the env-invalidation
// guarantee described in the tech framing.
//
// The canonical form is a sorted-key JSON object; sorted keys ensure that map
// iteration order differences never lead to different hashes for identical data.
func SchemeHash(s config.SecurityScheme) (string, error) {
	canonical, err := canonicalise(s)
	if err != nil {
		return "", fmt.Errorf("security hash: canonicalise failed: %w", err)
	}
	sum := sha256.Sum256(canonical)
	return hex.EncodeToString(sum[:]), nil
}

// canonicalise converts a SecurityScheme into a deterministic JSON byte slice.
// We marshal to a map first so we can sort the keys, then marshal back to JSON.
func canonicalise(s config.SecurityScheme) ([]byte, error) {
	// Marshal to an intermediate map so we get explicit key→value pairs.
	raw, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	// Unmarshal into a generic ordered representation.
	var m map[string]any
	decodeErr := json.Unmarshal(raw, &m)
	if decodeErr != nil {
		return nil, decodeErr
	}

	// Re-marshal with sorted keys.
	return marshalSorted(m)
}

// marshalSorted produces a JSON byte slice with all object keys sorted at every
// depth of the structure.  This guarantees identical output for equivalent data
// regardless of Go's map-iteration randomisation.
func marshalSorted(v any) ([]byte, error) {
	switch val := v.(type) {
	case map[string]any:
		return marshalSortedMap(val)

	case []any:
		return marshalSortedSlice(val)

	default:
		return json.Marshal(v)
	}
}

func marshalSortedMap(val map[string]any) ([]byte, error) {
	keys := make([]string, 0, len(val))
	for k := range val {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	buf := []byte{'{'}
	for i, k := range keys {
		key, err := json.Marshal(k)
		if err != nil {
			return nil, err
		}
		child, err := marshalSorted(val[k])
		if err != nil {
			return nil, err
		}
		buf = append(buf, key...)
		buf = append(buf, ':')
		buf = append(buf, child...)
		if i < len(keys)-1 {
			buf = append(buf, ',')
		}
	}
	return append(buf, '}'), nil
}

func marshalSortedSlice(val []any) ([]byte, error) {
	buf := []byte{'['}
	for i, item := range val {
		child, err := marshalSorted(item)
		if err != nil {
			return nil, err
		}
		buf = append(buf, child...)
		if i < len(val)-1 {
			buf = append(buf, ',')
		}
	}
	return append(buf, ']'), nil
}
