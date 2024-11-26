package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	files := []string{
		"form",
		"table",
		"visual",
	}

	for _, file := range files {
		http.HandleFunc(fmt.Sprintf("/%s", file), func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, fmt.Sprintf("e2e/server/%s.html", file))
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
