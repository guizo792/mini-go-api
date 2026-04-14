package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/guizo792/mini-go-api/internal/middleware"
)

func Handler(r *chi.Mux) {
	// Global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/user", func(router chi.Router) {
		// Middleware for /order routes
		router.Use(middleware.Authorization)
		router.Use(middleware.Recovery)
		router.Use(middleware.Logging)
		router.Get("/orders", GetOrder)
	})
}
