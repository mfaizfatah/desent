package handler

import (
	"encoding/json"
	"net/http"

	"desent/port"
)

type AuthHandler struct {
	authSvc port.AuthService
}

func (h *AuthHandler) Token(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	token, err := h.authSvc.GenerateToken(creds.Username, creds.Password)
	if err != nil {
		writeError(w, mapDomainError(err), err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"token": token})
}
