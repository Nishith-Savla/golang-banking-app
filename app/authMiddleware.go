package app

import (
	"github.com/Nishith-Savla/golang-banking-app/domain"
	"github.com/Nishith-Savla/golang-banking-app/errs"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	repo domain.AuthRepository
}

func (m AuthMiddleware) authorizationHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			currentRoute := mux.CurrentRoute(r)
			currentRouteVars := mux.Vars(r)
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				appError := errs.NewAuthorizationError("missing token")
				writeJSONResponse(w, appError.Code, appError.AsMessage())
				return
			}
			token := getTokenFromHeader(authHeader)

			isAuthorized := m.repo.IsAuthorized(token, currentRoute.GetName(), currentRouteVars)

			if isAuthorized {
				next.ServeHTTP(w, r)
			} else {
				appError := errs.NewForbiddenError("unauthorized")
				writeJSONResponse(w, appError.Code, appError.AsMessage())
			}
		})
	}
}

func getTokenFromHeader(header string) string {
	/*
	   token is coming in the format as below
	   "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50cyI6W.yI5NTQ3MCIsIjk1NDcyIiw"
	*/
	if strings.HasPrefix(header, "Bearer ") {
		return strings.TrimSpace(header[7:])
	}
	return ""
}
