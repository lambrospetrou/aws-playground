package main

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/coreos/go-systemd/v22/activation"
)

func main() {
	name := flag.String("name", "app", "The name of the service running, e.g. web2")
	port := flag.String("port", os.Getenv("PORT"), "The port for the server to listen.")
	isSystemd := flag.Bool("systemd", false, "Provide if the service should listen to the systemd socket activation listeners.")
	flag.Parse()

	if strings.TrimSpace(*port) == "" {
		*port = "5000"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("Service %s Path, %q", *name, html.EscapeString(r.URL.Path))
		log.Println(msg)
		fmt.Fprintf(w, msg)
	})

	if *isSystemd {

		path, err := os.Getwd()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(path) // for example /home/user
		path, err = os.Executable()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(path) // for example /tmp/go-build872132473/b001/exe/main

		// artificial delay to cause slow flip to test race condition with systemd and EB process...
		// time.Sleep(time.Second * time.Duration(10))

		listeners, err := activation.Listeners()
		if err != nil {
			panic(err)
		}
		if len(listeners) != 1 {
			panic("Unexpected number of socket activation fds")
		}
		log.Printf("BOOM8: App %s starts listening at the systemd file descriptor: %s\n", *name, listeners[0])
		log.Fatalln(http.Serve(listeners[0], nil))
	} else {
		log.Printf("App %s starts listening at :%s\n", *name, *port)
		log.Fatalln(http.ListenAndServe(":"+*port, nil))
	}
}
