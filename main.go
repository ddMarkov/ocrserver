package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func main() {

	r := chi.NewRouter()
	r.Post("/", ExampleHandler)

	port := "8080"
	if envPort := os.Getenv("PORT"); len(envPort) > 0 {
		port = envPort
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), r); err != nil {
		fmt.Println(err)
	}
}

// ExampleHandler .
func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	w.Write([]byte("hello world"))
}
