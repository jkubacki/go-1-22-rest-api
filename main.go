package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("http://localhost:8080")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /comment", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "return all comments")
	})

	mux.HandleFunc("GET /comment/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "return comment %s", id)
	})

	mux.HandleFunc("POST /comment", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "post a new comment")
	})

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println(err.Error())
	}
}
