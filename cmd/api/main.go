package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"desent/internal/adapter/handler"
	"desent/internal/adapter/repository"
	"desent/internal/service"
)

func main() {
	bookRepo := repository.NewMemoryBookRepository()
	tokenRepo := repository.NewMemoryTokenRepository(1 * time.Hour)

	bookSvc := service.NewBookService(bookRepo)
	authSvc := service.NewAuthService(tokenRepo)

	router := handler.NewRouter(bookSvc, authSvc)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
