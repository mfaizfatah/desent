package port

import "desent/domain"

type BookService interface {
	Create(book domain.Book) (domain.Book, error)
	GetByID(id string) (domain.Book, error)
	List(filter domain.BookFilter) (any, error)
	Update(id string, book domain.Book) (domain.Book, error)
	Delete(id string) error
}

type AuthService interface {
	GenerateToken(username, password string) (string, error)
	ValidateToken(token string) error
}
