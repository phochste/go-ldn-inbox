package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: ldn-consumer url")
		return
	}

	url := os.Args[1]

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: received non-200 status code:", resp.StatusCode)
		os.Exit(1)
	}

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		os.Exit(1)
	}

	context, exists := result["@context"]

	if exists && context == "http://www.w3.org/ns/ldp" {
		contains, _ := result["contains"].([]interface{})
		for _, item := range contains {
			fmt.Println(item)
		}
	} else {
		jsonData, err := json.MarshalIndent(result, "", "  ")

		if err != nil {
			fmt.Println("Error marshaling to JSON:", err)
			return
		}

		fmt.Println(string(jsonData))
	}
}
