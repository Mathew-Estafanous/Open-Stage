package middleware

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// CORS creates a middleware that will handle CORS for the API.
func CORS() mux.MiddlewareFunc {
	headerOpts := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Set-Cookies"})
	originOpts := handlers.AllowedOrigins([]string{"http://localhost:3000", "https://open-stage-web.herokuapp.com/"})
	methodOpts := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	credentialsOpts := handlers.AllowCredentials()
	return handlers.CORS(headerOpts, originOpts, methodOpts, credentialsOpts)
}
