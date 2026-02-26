package repository

import (
	"fmt"
	"strconv"
	"sync"

	"desent/domain"
)

type MemoryBookRepository struct {
	mu     sync.RWMutex
	books  map[string]domain.Book
	nextID int
}

func NewMemoryBookRepository() *MemoryBookRepository {
	return &MemoryBookRepository{
		books:  make(map[string]domain.Book),
		nextID: 1,
	}
}

func (r *MemoryBookRepository) Save(book domain.Book) domain.Book {
	r.mu.Lock()
	defer r.mu.Unlock()
	book.ID = strconv.Itoa(r.nextID)
	r.nextID++
	r.books[book.ID] = book
	return book
}

func (r *MemoryBookRepository) FindByID(id string) (domain.Book, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	book, ok := r.books[id]
	if !ok {
		return domain.Book{}, fmt.Errorf("book with id '%s' not found", id)
	}
	return book, nil
}

func (r *MemoryBookRepository) FindAll() []domain.Book {
	r.mu.RLock()
	defer r.mu.RUnlock()
	books := make([]domain.Book, 0, len(r.books))
	for _, b := range r.books {
		books = append(books, b)
	}
	return books
}

func (r *MemoryBookRepository) Update(id string, book domain.Book) (domain.Book, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.books[id]; !ok {
		return domain.Book{}, fmt.Errorf("book with id '%s' not found", id)
	}
	book.ID = id
	r.books[id] = book
	return book, nil
}

func (r *MemoryBookRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.books[id]; !ok {
		return fmt.Errorf("book with id '%s' not found", id)
	}
	delete(r.books, id)
	return nil
}
