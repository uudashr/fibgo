package main

import (
	"log"
	"net/http"

	fib "github.com/uudashr/fibgo"
)

func main() {
	handler := fib.NewHTTPHandler()
	log.Println("Listening on port", 8080, "...")
	err := http.ListenAndServe(":8080", handler)
	log.Println("Stopped err:", err)
}
