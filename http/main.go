// input url, response status code and response body as well.
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Usage: go run main.go <url>")
		os.Exit(1)
	}

	if _, err := url.ParseRequestURI(args[1]); err != nil {
		fmt.Println("Invalid URL")
		os.Exit(1)
	}

	resp, err := http.Get(args[1])

	if err != nil {
		fmt.Println("Error fetching URL")
		os.Exit(1)
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode
	fmt.Println("Status Code:", statusCode)

	if statusCode == 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body")
			os.Exit(1)
		}
		fmt.Println("Response Body:")
		fmt.Printf("%s", body)
	}

}
