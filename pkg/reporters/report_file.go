package reporters

import (
	"log"
	"os"
	"path/filepath"
	"testflowkit/internal/utils/fileutils"
)

func writeReportFile(path string, data []byte) {
	if mkdirErr := os.MkdirAll(filepath.Dir(path), fileutils.DirPermission); mkdirErr != nil {
		log.Panicf("cannot create report directory ( %s )\n", mkdirErr)
	}

	if writeErr := os.WriteFile(path, data, fileutils.FilePermission); writeErr != nil {
		log.Panicf("cannot write report file %s ( %s )\n", path, writeErr)
	}
}

// groupByFeature groups regular scenarios by feature URI, preserving first-seen order. Hooks are excluded.
func groupByFeature(scenarios []Scenario) ([]string, map[string][]Scenario) {
	var order []string
	groups := map[string][]Scenario{}
	for _, sc := range scenarios {
		if sc.Type != scenario {
			continue
		}
		if _, ok := groups[sc.URI]; !ok {
			order = append(order, sc.URI)
		}
		groups[sc.URI] = append(groups[sc.URI], sc)
	}
	return order, groups
}
