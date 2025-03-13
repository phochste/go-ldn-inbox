package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() { os.Exit(mainReturnWithCode()) }

func mainReturnWithCode() int {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: ldn-consumer url")
		return 1
	}

	url := os.Args[1]

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return 2
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: received non-200 status code:", resp.StatusCode)
		return 3
	}

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return 4
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
			return 5
		}

		fmt.Println(string(jsonData))
	}

	return 0
}
