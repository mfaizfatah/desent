package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"desent/domain"
	"desent/port"
)

func NewRouter(bookSvc port.BookService, authSvc port.AuthService) http.Handler {
	mux := http.NewServeMux()

	ping := &PingHandler{}
	book := &BookHandler{bookSvc: bookSvc, authSvc: authSvc}
	auth := &AuthHandler{authSvc: authSvc}

	mux.HandleFunc("GET /ping", ping.Ping)
	mux.HandleFunc("POST /echo", ping.Echo)

	mux.HandleFunc("POST /auth/token", auth.Token)

	mux.HandleFunc("GET /books", book.List)
	mux.HandleFunc("GET /books/{id}", book.GetByID)
	mux.HandleFunc("POST /books", book.Create)
	mux.HandleFunc("PUT /books/{id}", book.Update)
	mux.HandleFunc("DELETE /books/{id}", book.Delete)

	return mux
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

func extractToken(r *http.Request) string {
	auth := r.Header.Get("Authorization")
	if strings.HasPrefix(auth, "Bearer ") {
		return strings.TrimPrefix(auth, "Bearer ")
	}
	return ""
}

func mapDomainError(err error) int {
	switch {
	case errors.Is(err, domain.ErrBookNotFound) || strings.Contains(err.Error(), "not found"):
		return http.StatusNotFound
	case errors.Is(err, domain.ErrUnauthorized):
		return http.StatusUnauthorized
	case errors.Is(err, domain.ErrTitleRequired),
		errors.Is(err, domain.ErrAuthorRequired),
		errors.Is(err, domain.ErrYearRequired),
		errors.Is(err, domain.ErrInvalidJSON),
		errors.Is(err, domain.ErrCredentials):
		return http.StatusBadRequest
	case errors.Is(err, domain.ErrInvalidCredentials):
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
