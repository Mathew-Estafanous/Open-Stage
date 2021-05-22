package middleware

import (
	"net/http"
	"os"
)

// EnforceSSL will redirect all HTTP requests to HTTPS to ensure SSL security.
// SSL is only enforced if the app is running in production. If not, regular http
// requests will be accepted.
func EnforceSSL(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if os.Getenv("PROFILE") == "prod" {
			if r.Header.Get("x-forwarded-proto") != "https" {
				sslUrl := "https://" + r.Host + r.RequestURI
				http.Redirect(w, r, sslUrl, http.StatusTemporaryRedirect)
				return
			}
		}

		h.ServeHTTP(w, r)
	})
}
