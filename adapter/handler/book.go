package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"desent/domain"
	"desent/port"
)

type BookHandler struct {
	bookSvc port.BookService
	authSvc port.AuthService
}

func (h *BookHandler) Create(w http.ResponseWriter, r *http.Request) {
	var book domain.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	created, err := h.bookSvc.Create(book)
	if err != nil {
		writeError(w, mapDomainError(err), err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, created)
}

func (h *BookHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	book, err := h.bookSvc.GetByID(id)
	if err != nil {
		writeError(w, mapDomainError(err), err.Error())
		return
	}
	writeJSON(w, http.StatusOK, book)
}

func (h *BookHandler) List(w http.ResponseWriter, r *http.Request) {
	if token := extractToken(r); token != "" {
		if err := h.authSvc.ValidateToken(token); err != nil {
			writeError(w, http.StatusUnauthorized, err.Error())
			return
		}
	}

	filter := domain.BookFilter{
		Author: r.URL.Query().Get("author"),
	}
	if p := r.URL.Query().Get("page"); p != "" {
		filter.Page, _ = strconv.Atoi(p)
	}
	if l := r.URL.Query().Get("limit"); l != "" {
		filter.Limit, _ = strconv.Atoi(l)
	}

	result, err := h.bookSvc.List(filter)
	if err != nil {
		writeError(w, mapDomainError(err), err.Error())
		return
	}

	writeJSON(w, http.StatusOK, result)
}

func (h *BookHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var book domain.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	updated, err := h.bookSvc.Update(id, book)
	if err != nil {
		writeError(w, mapDomainError(err), err.Error())
		return
	}

	writeJSON(w, http.StatusOK, updated)
}

func (h *BookHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := h.bookSvc.Delete(id); err != nil {
		writeError(w, mapDomainError(err), err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
