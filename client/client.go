package main

import (
	"io"
	"context"
	"fmt"
	"net/http"
	"time"
	"os"
	"bytes"
	"mime/multipart"
)

func main() {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	
	file, err := os.Open("test.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("test", "test.txt")
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		panic(err)
	}
	
	req, err := http.NewRequestWithContext(context.Background(), "POST", "http://localhost:8080/", body)
	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("unexpected status: got %v", res.Status))
	}
	fmt.Println(res.Header.Get("Content-Type"))
	b, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
