/*
	An HTTP server for mocking responses.

	Installation:

	1. $ go build httpm.go
	2. Put the file in a directory in your $PATH.

	Usage:

	$ httpm <flags>

	Flags:

	-port: The port to run the webserver on. Default 8000.
	-delay: The delay in seconds before responding. Default 0.
	-status: The HTTP status code of the response. Default 200.
	-content-type: The content type of the response. Default "text/plain".
	-filename: File to read for the response body. Default "".
*/

package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	delay       int
	statusCode  int
	contentType string
	fileName    string
)

func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(delay) * time.Second)

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)

	if len(fileName) == 0 {
		io.WriteString(w, "")
		return
	}

	body, err := ioutil.ReadFile(fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	io.WriteString(w, string(body))
}

func main() {
	port := flag.Int("port", 8000, "The server port.")

	flag.IntVar(&delay, "delay", 0, "Response delay.")
	flag.IntVar(&statusCode, "status-code", 200, "Response status code.")
	flag.StringVar(&contentType, "content-type", "text/plain",
		"Response content type.")
	flag.StringVar(&fileName, "filename", "", "Response body filename.")

	flag.Parse()

	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(": %v", *port), nil)
}
