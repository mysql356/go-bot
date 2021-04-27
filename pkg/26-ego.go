package pkg

import (
	"log"
	"net/http"
)

func Ego() {
	// Simple static webserver:
	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("/Library/WebServer/Documents/php_test"))))
	mux.Handle("/uploads/", http.StripPrefix("/uploads", http.FileServer(http.Dir("/Library/WebServer/Documents/php_test/ego/uploads/"))))

	log.Fatal(http.ListenAndServe(":3000", mux))
}
