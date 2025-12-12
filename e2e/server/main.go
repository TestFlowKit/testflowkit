package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"runtime"
	"time"
)

func main() {
	files := []string{
		"form",
		"table",
		"visual",
		"details",
		"scroll",
		"file-upload",
	}

	_, filename, _, _ := runtime.Caller(0)
	currDir := filepath.Dir(filename)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filePath := path.Join(currDir, "index.html")
		http.ServeFile(w, r, filePath)
	})

	for _, file := range files {
		http.HandleFunc("/"+file, func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.URL.Path, "requested")
			filePath := path.Join(currDir, file+".html")
			http.ServeFile(w, r, filePath)
		})
	}

	const port = 3000
	const timeout = 3
	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		ReadHeaderTimeout: timeout * time.Second,
	}

	log.Printf("tests e2e server launched on port %d\n", port)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
