package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// my personal http client

func main() {
	client := &http.Client{}

	resp, err := client.Get("https://jw.org")
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
