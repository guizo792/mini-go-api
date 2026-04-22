package handlers

import (
	"time"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/guizo792/mini-go-api/internal/middleware"
	"github.com/guizo792/mini-go-api/internal/tools"
)

func Handler(r *chi.Mux) error {
	db, err := tools.NewDatabase(false)
	if err != nil {
		return err
	}

	orderHandler := OrderHandler{DB: db}

	// Global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/user", func(router chi.Router) {
		// Middleware for /order routes
		router.Use(middleware.RateLimit(5, time.Minute))
		router.Use(middleware.Authorization)
		router.Use(middleware.Recovery)
		router.Use(middleware.Logging)
		router.Get("/orders", orderHandler.GetOrder)
	})

	return nil
}
