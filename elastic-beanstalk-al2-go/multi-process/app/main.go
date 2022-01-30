package main

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	name := flag.String("name", "app", "The name of the service running, e.g. web2")
	port := flag.String("port", os.Getenv("PORT"), "The port for the server to listen.")
	flag.Parse()

	if strings.TrimSpace(*port) == "" {
		*port = "5000"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("Service %s Path, %q", *name, html.EscapeString(r.URL.Path))
		log.Println(msg)
		fmt.Fprintf(w, msg)
	})

	log.Printf("App %s starts listening at :%s\n", *name, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
