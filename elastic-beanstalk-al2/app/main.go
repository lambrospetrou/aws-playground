package main

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	name := flag.String("name", "app1", "The name of the service running, e.g. web2")
	port := flag.String("port", "5000", "The port for the server to listen.")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Service %s Path, %q", *name, html.EscapeString(r.URL.Path))
	})

	log.Printf("App %s starts listening at :%s\n", *name, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
