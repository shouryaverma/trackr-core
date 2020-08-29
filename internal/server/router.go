package server

import (
	"github.com/amaraliou/trackr-v2/internal/handler"
	trackrMiddleware "github.com/amaraliou/trackr-v2/internal/middleware"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

// NewRouter ...
func (server *Server) NewRouter(handler *handler.Handler) error {
	router := chi.NewRouter()

	// Cors
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	// Middlewares
	router.Use(cors.Handler)
	router.Use(middleware.StripSlashes)
	router.Use(trackrMiddleware.SetJSON)
	router.Use(middleware.Heartbeat("/ping"))

	router.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/login", handler.Login)

		r.Post("/users", handler.CreateUser)
		r.Get("/users", handler.GetAllUsers)
		r.With(trackrMiddleware.SetAuth).Get("/users/{id}", handler.GetUser)
		r.With(trackrMiddleware.SetAuth).Put("/users/{id}", handler.UpdateUser)
		r.With(trackrMiddleware.SetAuth).Delete("/users/{id}", handler.DeleteUser)
	})

	server.Router = router
	return nil
}
