package main

import (
	"fmt"
	"net/http"
)

// ハンドラー関数を定義
func hello(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Hello!")
}

func world(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "World!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	server.ListenAndServe()
}
