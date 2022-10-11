package main

import (
	"log"
	"net/http"
	"os"

	"com.lambrospetrou/aws-playground/aws-lambda-url-and-fly-golang-proxy/core"
)

func main() {
	mutex := core.NewMutex()
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "5000"
	}
	if err := http.ListenAndServe(":"+port, mutex); err != nil {
		log.Fatal(err)
	}
}
