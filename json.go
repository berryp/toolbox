package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

func getFileContents(path string) (string, error) {
	// Ensure the file exists before trying to open it.
	if _, err := os.Stat(path); os.IsNotExist(err) {
		msg := fmt.Sprintf("No such file or directory: %s", path)
		return "", errors.New(msg)
	}

	// Get the contents of the file.
	data, err := ioutil.ReadFile(path)
	if err != nil {
		msg := fmt.Sprintf("Error opening file: %s", path)
		return "", errors.New(msg)
	}

	return string(data), nil
}

func getURLContents(path string) (string, error) {
	// Fetch the content from the given URL.
	resp, err := http.Get(path)
	if err != nil {
		msg := fmt.Sprintf("Error fetching content: %s", path)
		return "", errors.New(msg)
	}

	// Make sure the request returns JSON content type.
	contentType := resp.Header["Content-Type"][0]
	re := regexp.MustCompile("application/json")
	result := re.FindString(contentType)

	if result != "application/json" {
		msg := fmt.Sprintf("Unsupported content type: %s", contentType)
		return "", errors.New(msg)
	}

	defer resp.Body.Close()

	// Read the response body.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		msg := fmt.Sprintf("Error reading content: %s", path)
		return "", errors.New(msg)
	}

	return string(body), nil
}

func readPath(path string) (string, error) {
	re := regexp.MustCompile("https?://")
	result := re.FindStringIndex(path)

	// If path is a URL then fetch the contents of the URL.
	if len(result) > 0 && result[0] == 0 {
		return getURLContents(path)
	}

	return getFileContents(path)
}

func usage() {
	fmt.Println("Usage: json [path]")
}

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		usage()
		os.Exit(1)
	}

	// Try to read the path.
	path := args[0]
	content, err := readPath(path)
	if err != nil {
		usage()
		fmt.Printf("\nError: %s\n", err)
		os.Exit(1)
	}

	var doc interface{}

	// Ensure the document is a valid JSON document.
	err = json.Unmarshal([]byte(content), &doc)
	if err != nil {
		usage()
		fmt.Printf("\nError: %s\n", err)
		os.Exit(1)
	}

	// Finally, pretty print the JSON document.
	b, _ := json.MarshalIndent(doc, "", "  ")
	fmt.Println(string(b))
}
