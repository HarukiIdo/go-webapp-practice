package main

import (
	"log"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
