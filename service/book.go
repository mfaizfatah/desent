package service

import (
	"sort"
	"strconv"
	"strings"

	"desent/domain"
	"desent/port"
)

type BookService struct {
	repo port.BookRepository
}

func NewBookService(repo port.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) Create(book domain.Book) (domain.Book, error) {
	if err := book.Validate(); err != nil {
		return domain.Book{}, err
	}
	return s.repo.Save(book), nil
}

func (s *BookService) GetByID(id string) (domain.Book, error) {
	return s.repo.FindByID(id)
}

func (s *BookService) List(filter domain.BookFilter) (any, error) {
	books := s.repo.FindAll()

	if filter.Author != "" {
		filtered := make([]domain.Book, 0)
		for _, b := range books {
			if strings.EqualFold(b.Author, filter.Author) {
				filtered = append(filtered, b)
			}
		}
		books = filtered
	}

	sort.Slice(books, func(i, j int) bool {
		idI, _ := strconv.Atoi(books[i].ID)
		idJ, _ := strconv.Atoi(books[j].ID)
		return idI < idJ
	})

	if filter.Page > 0 || filter.Limit > 0 {
		page := filter.Page
		limit := filter.Limit
		if page < 1 {
			page = 1
		}
		if limit < 1 {
			limit = 10
		}

		total := len(books)
		start := (page - 1) * limit
		end := start + limit
		if start > total {
			start = total
		}
		if end > total {
			end = total
		}

		return books[start:end], nil
	}

	return books, nil
}

func (s *BookService) Update(id string, book domain.Book) (domain.Book, error) {
	if err := book.Validate(); err != nil {
		return domain.Book{}, err
	}
	return s.repo.Update(id, book)
}

func (s *BookService) Delete(id string) error {
	return s.repo.Delete(id)
}
