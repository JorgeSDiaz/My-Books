package server

import (
	"net/http"

	healthHandler "github.com/JorgeSDiaz/My-Books/cmd/server/handlers/health"

	"github.com/JorgeSDiaz/My-Books/internal/book"
	"github.com/JorgeSDiaz/My-Books/internal/health"
)

func routes() {
	healthHandler := healthHandler.NewHandler(health.NewService())

	http.HandleFunc("/health", healthHandler.Health)
}
