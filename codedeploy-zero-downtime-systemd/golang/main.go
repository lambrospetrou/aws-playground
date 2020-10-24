package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/coreos/go-systemd/v22/activation"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "Hello World! Append a name to the URL to say hello. For example, use %s/Mary to say hello to Mary.", r.Host)
	} else {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	}
}

func main() {
	port := os.Getenv("PORT")
	startedFromSystemd := os.Getenv("STARTED_FROM_SYSTEMD") == "1"

	listeners, err := activation.Listeners()
	if err != nil && startedFromSystemd {
		panic(err)
	}

	http.HandleFunc("/", handler)

	if startedFromSystemd {
		if len(listeners) != 1 {
			panic("Unexpected number of socket activation fds")
		}
		log.Println("Starting the server listening to the systemd file descriptor...", listeners[0])
		http.Serve(listeners[0], nil)
	} else {
		if port == "" {
			port = "0"
		}
		log.Println("Starting the server listening to port:", port)
		http.ListenAndServe(":"+port, nil)
	}
}
