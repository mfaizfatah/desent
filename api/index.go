package handler

import (
	"net/http"
	"time"

	httphandler "desent/internal/adapter/handler"
	"desent/internal/adapter/repository"
	"desent/internal/service"
)

var router http.Handler

func init() {
	bookRepo := repository.NewMemoryBookRepository()
	tokenRepo := repository.NewMemoryTokenRepository(1 * time.Hour)

	bookSvc := service.NewBookService(bookRepo)
	authSvc := service.NewAuthService(tokenRepo)

	router = httphandler.NewRouter(bookSvc, authSvc)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
