package main

import (
	"log"
	"net/http"

	"tech-test/internal/domain"
	"tech-test/internal/handlers"
)

func main() {
	// Initialize domain services
	mathService := domain.NewMathService()

	// Initialize handlers with dependency injection
	h := handlers.NewHandlers(mathService)

	// Setup routes
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", h.Ping)
	mux.HandleFunc("/add", h.Add)
	mux.HandleFunc("/sub", h.Sub)
	mux.HandleFunc("/mul", h.Mul)

	// Start server on port 8080
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
