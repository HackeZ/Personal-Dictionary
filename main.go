package main

import (
	"log"
	"net/http"

	ctrl "Personal-Dictionary/src/back-end/controllers"

	"github.com/bmizerany/pat"
)

func main() {
	log.Printf("=== Personal Dictionary Server Start ===")

	// == RESTful API Router ==
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(ctrl.Index))

	mux.Get("/login", http.HandlerFunc(ctrl.Login))

	// == RESTful API Router
}
