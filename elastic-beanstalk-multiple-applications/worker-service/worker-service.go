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
		port = "8083"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Worker Path, %q", html.EscapeString(r.URL.Path))
	})

	log.Println("Worker Service starts listening at :" + port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
