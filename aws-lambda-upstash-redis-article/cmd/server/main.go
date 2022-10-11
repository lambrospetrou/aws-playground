package main

import (
	"log"
	"net/http"
	"os"

	"com.lambrospetrou/aws-playground/aws-lambda-upstash-redis-article/core"
)

func main() {
	mux := core.NewMux()

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "5000"
	}
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
