package port

import "desent/internal/domain"

type BookRepository interface {
	Save(book domain.Book) domain.Book
	FindByID(id string) (domain.Book, error)
	FindAll() []domain.Book
	Update(id string, book domain.Book) (domain.Book, error)
	Delete(id string) error
}

type TokenRepository interface {
	Store(token string)
	Exists(token string) bool
}
