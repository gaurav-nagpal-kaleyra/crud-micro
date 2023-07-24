package middleware

import (
	"context"
	"net/http"

	"go.uber.org/zap"
)

func AddAuthKey(next http.Handler) http.Handler {
	zap.L().Debug("Middleware called")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "authenticated", true)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
