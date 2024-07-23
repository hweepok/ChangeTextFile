package main

import (
	"fmt"
	"net/http"
	"log"
	"io"
	"strings"
)

func handle(rw http.ResponseWriter, r *http.Request) {
	log.Println("Received request")
	text, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	
	newText := strings.ReplaceAll(string(text), "this", "butt")
	rw.Write([]byte(newText))
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", handle)

	server := http.Server{
		Addr: ":8080",
		Handler: router,
	}

	fmt.Println("Server listening on port :8080")
	err := server.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}
