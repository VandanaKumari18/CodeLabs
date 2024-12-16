package main

import (
	"fmt"
	"net/http"

	"CodeLabs/internal/user/api"
	"github.com/go-chi/chi/v5"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, Welcome to CodeLabs!")
	})

	r := chi.NewRouter()

	// routes for the User Module
	r.Post("/signup", api.SignupHandler)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
