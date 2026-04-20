package shared

import (
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"testflowkit/internal/step_definitions/backend"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/internal/step_definitions/frontend"
	"testflowkit/internal/step_definitions/frontend/assertions"
	"testflowkit/internal/step_definitions/variables"
	"testflowkit/internal/utils/fileutils"
)

var sentenceCleanRe = regexp.MustCompile(`[$^]`)

// GetAllDocs returns documentation for all registered step definitions.
func GetAllDocs() []stepbuilder.Documentation {
	return slices.Concat(frontend.GetDocs(), backend.GetDocs(), variables.GetDocs(), assertions.GetDocs())
}

// CleanSentence strips regex anchors from a sentence string.
func CleanSentence(sentence string) string {
	return sentenceCleanRe.ReplaceAllString(sentence, "")
}

// WriteFileWithDirectories creates all parent directories then writes data to filePath.
func WriteFileWithDirectories(filePath string, data []byte) error {
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, fileutils.DirPermission); err != nil {
		return err
	}
	return os.WriteFile(filePath, data, fileutils.FilePermission)
}
