package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	fib "github.com/uudashr/fibgo"
)

var usageMsg = `
Usage: fibgo-server [options]
 -p, --port <port>    Listen port
 -h, --help           Show this message
`

func usage() {
	fmt.Println(usageMsg)
}

func main() {
	var port int
	var showHelp bool

	flag.IntVar(&port, "port", 8080, "Listen port")
	flag.IntVar(&port, "p", 8080, "Listen port")
	flag.BoolVar(&showHelp, "help", false, "Show this message")
	flag.BoolVar(&showHelp, "h", false, "Show this message")

	flag.Usage = usage
	flag.Parse()

	if showHelp {
		flag.Usage()
		return
	}

	handler := fib.NewHTTPHandler()
	log.Println("Listening on port", port, "...")
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
	log.Println("Stopped err:", err)
}
