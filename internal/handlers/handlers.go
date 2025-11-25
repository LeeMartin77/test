package handlers

import (
	"fmt"
	"net/http"
	"tech-test/internal/domain"
)

type Handlers struct {
	mathService domain.MathService
}

func NewHandlers(mathService domain.MathService) *Handlers {
	return &Handlers{
		mathService: mathService,
	}
}

func (h *Handlers) Ping(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "pong")
}
