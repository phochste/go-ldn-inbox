package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "usage: ldn-sender url file")
		return
	}

	postUrl := os.Args[1]
	postFile := os.Args[2]

	postContent, err := os.ReadFile(postFile)

	if err != nil {
		fmt.Fprintf(os.Stderr, "no such file "+postFile)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", postUrl, bytes.NewBuffer(postContent))

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/ld+json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Print the response status
	fmt.Println("Response status:", resp.Status)
}
