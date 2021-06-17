package middleware

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// CORS is creates a middleware that will handle CORS for the API.
func CORS() mux.MiddlewareFunc {
	headerOpts := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originOpts := handlers.AllowedOrigins([]string{"*"})
	methodOpts := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	return handlers.CORS(headerOpts, originOpts, methodOpts)
}
