package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Words struct {
	Page  string   `json:"page"`
	Input string   `json:"input"`
	Words []string `json:"words"`
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: go run main.go <url>")
		return
	}

	if _, err := http.Get(args[1]); err != nil {
		fmt.Println("Error: ", err)
		return
	}

	resp, err := http.Get(args[1])

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	var words Words
	err = json.Unmarshal(body, &words)
	if err != nil {
		fmt.Println("Error unmarshalling JSON")
		return
	}

	fmt.Printf("Page: %s\nWords: %s\n", words.Page, words.Words)
}
