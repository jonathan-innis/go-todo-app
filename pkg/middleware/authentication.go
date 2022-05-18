package middleware

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jonathan-innis/go-todo-app/pkg/auth"
	"github.com/jonathan-innis/go-todo-app/pkg/helper"
)

const (
	AuthorizationKey = "Authorization"
	UserIdKey        = "userId"
)

func UserAuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: validate the bearer token for this user
		vars := mux.Vars(r)
		userId, ok := vars[UserIdKey]
		if !ok {
			helper.GetError(w, http.StatusNotFound, "User Id not found in the request")
			return
		}
		authHeader := r.Header.Get(AuthorizationKey)
		if !ok {
			helper.GetError(w, http.StatusUnauthorized, "Authorization header not included in the request")
			return
		}
		err := auth.ValidateToken(userId, auth.ParseToken(authHeader))
		if err != nil {
			helper.GetError(w, http.StatusUnauthorized, fmt.Sprintf("Token is not authorized to view user %v", userId))
			return
		}
		next.ServeHTTP(w, r)
	})
}
