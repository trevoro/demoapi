package middleware

import "net/http"

type AuthMiddleware struct{}

func (m *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	next(w, r)
}
