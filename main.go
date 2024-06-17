package main

import (
	"api/middleware"
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	router.HandleFunc("GET /comment", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "return all comments")
	})

	router.HandleFunc("GET /comment/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "return comment %s", id)
	})

	router.HandleFunc("POST /comment", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "post a new comment")
	})

	stack := middleware.CreateStack(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: stack(router),
	}

	fmt.Println("Server listening at http://localhost:8080")

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err.Error())
	}
}
