package testsconfig

import (
	"fmt"
	"os"
	"path/filepath"
	"net/http"
)

var pages = map[string]string{
	"Test": "http://localhost:8080/test.html",
}

func GetPageURL(page string) (string, error) {
	url, exists := pages[page]
	if !exists {
		return "", fmt.Errorf("page %s not found", page)
	}
	return url, nil
}

func IsPageDefined(page string) bool {
	_, exists := pages[page]
	return exists
}

func init() {
	// Start a simple HTTP server to serve the test page
	go func() {
		http.HandleFunc("/test.html", func(w http.ResponseWriter, r *http.Request) {
			content, err := os.ReadFile(filepath.Join("e2e", "server", "test.html"))
			if err != nil {
				http.Error(w, "Error reading file", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "text/html")
			w.Write(content)
		})
		http.ListenAndServe(":8080", nil)
	}()
} 