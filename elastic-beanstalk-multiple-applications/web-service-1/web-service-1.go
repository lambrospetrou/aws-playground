package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Service 1 Path, %q", html.EscapeString(r.URL.Path))
	})

	log.Println("Web Service 1 starts listening at :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
