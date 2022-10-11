package core

import (
	"embed"
	"io"
	"log"
	"net/http"
)

//go:embed files/*
var Files embed.FS

func NewMutex() *http.ServeMux {
	mutex := http.NewServeMux()
	mutex.HandleFunc("/xxx/files/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		log.Println(r.Header.Get("X_XXX_CDN_AUTH"))
		http.StripPrefix("/xxx", http.FileServer(http.FS(Files))).ServeHTTP(w, r)
		log.Println("Files delivered!")
	})
	mutex.HandleFunc("/files/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		log.Println(r.Header.Get("X_XXX_CDN_AUTH"))
		http.FileServer(http.FS(Files)).ServeHTTP(w, r)
		log.Println("Files delivered!")
	})
	mutex.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		log.Println(r.Header.Get("X_XXX_CDN_AUTH"))
		io.WriteString(w, "Hello")
	})
	return mutex
}
