package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func main() { os.Exit(mainReturnWithCode()) }

func mainReturnWithCode() int {

	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "usage: ldn-sender url file")
		return 1
	}

	postUrl := os.Args[1]
	postFile := os.Args[2]

	postContent, err := os.ReadFile(postFile)

	if err != nil {
		fmt.Fprintf(os.Stderr, "no such file "+postFile)
		return 2
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", postUrl, bytes.NewBuffer(postContent))

	if err != nil {
		fmt.Println(err)
		return 3
	}

	req.Header.Set("Content-Type", "application/ld+json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return 4
	}
	defer resp.Body.Close()

	// Print the response status
	fmt.Println("Response status:", resp.Status)

	return 0
}
