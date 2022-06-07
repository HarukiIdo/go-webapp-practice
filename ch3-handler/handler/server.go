package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello!!!")
}

func main() {
	helloHandler := HelloHandler{}

	server := http.Server{
		Addr:    ":8080",
		Handler: &helloHandler,
	}

	server.ListenAndServe()

}
