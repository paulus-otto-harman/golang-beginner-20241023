package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	parentCtx := context.Background()

	ctx, cancel := context.WithTimeout(parentCtx, 3*time.Second)
	defer cancel()

	url := "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Layanan tersedia")
	} else {
		fmt.Println("Layanan tidak tersedia")
	}
}
