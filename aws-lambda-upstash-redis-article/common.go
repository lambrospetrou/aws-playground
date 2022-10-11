package main

import (
	"log"
	"time"
)

func timeit(name string, f func()) {
	start := time.Now()
	f()
	duration := time.Since(start)

	log.Println(name+" elapsed time: ", duration)
}
