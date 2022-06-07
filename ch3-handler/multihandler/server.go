package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}
type WorldHandler struct{}

func (h *HelloHandler) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(wr, "Hello!")
}

func (h *WorldHandler) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(wr, "world!")
}
func main() {
	hello := HelloHandler{}
	world := WorldHandler{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.Handle("/hello/", &hello)
	http.Handle("/world/", &world)

	server.ListenAndServe()
}
