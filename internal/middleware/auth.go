package middleware

import (
	"net/http"

	"github.com/amaraliou/trackr-core/internal/auth"
	"github.com/amaraliou/trackr-core/internal/response"
)

// SetAuth ...
func SetAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		err := auth.TokenValid(request)
		if err != nil {
			response.ERROR(writer, http.StatusUnauthorized, err)
			return
		}
		next.ServeHTTP(writer, request)
	})
}
