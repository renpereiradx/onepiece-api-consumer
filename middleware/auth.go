package middleware

import (
	"net/http"
	"strings"

	"github.com/renpereiradx/onepiece-api-consumer/helpers"
	"github.com/renpereiradx/onepiece-api-consumer/server"
)

var NO_AUTH_NEEDED = []string{
	"login",
	"signup",
}

func shouldCheckToken(route string) bool {
	for _, p := range NO_AUTH_NEEDED {
		if strings.Contains(route, p) {
			return false
		}
	}
	return true
}

func CheckAuthMiddleware(server server.Server) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !shouldCheckToken(r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}
			_, err := helpers.GetJWTAuthorizationInfo(server, w, r)
			if err != nil {
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
