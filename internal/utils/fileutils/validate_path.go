package fileutils

import (
	"fmt"
	"path/filepath"
	"strings"
)

func ValidatePath(path string, allowedPrefixes ...string) error {
	cleanPath := filepath.Clean(path)
	sep := string(filepath.Separator)

	if cleanPath == ".." || strings.HasPrefix(cleanPath, ".."+sep) || strings.Contains(cleanPath, sep+"..") {
		return fmt.Errorf("invalid path contains directory traversal: %s", path)
	}

	if len(allowedPrefixes) == 0 {
		return nil
	}

	for _, prefix := range allowedPrefixes {
		if cleanPath == prefix || strings.HasPrefix(cleanPath, prefix+string(filepath.Separator)) {
			return nil
		}
	}

	return fmt.Errorf("path %s is not within allowed locations: %v", path, allowedPrefixes)
}
