/*
	Statically serve file tree from the current directory.

	Installation:

	1. $ go build serve.go
	2. Put the file in a directory in your $PATH.

	Usage:

	$ serve [flags]

	Flags:

	-port: The port to run the web server on. Default: 5000.
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func LogHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	port := flag.Int("port", 5000, "")
	flag.Parse()

	fs := http.FileServer(http.Dir("."))
	http.Handle("/", LogHandler(fs))

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", *port), nil))
}
