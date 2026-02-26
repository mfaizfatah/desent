package domain

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func (b Book) Validate() error {
	if b.Title == "" {
		return ErrTitleRequired
	}
	if b.Author == "" {
		return ErrAuthorRequired
	}
	if b.Year == 0 {
		return ErrYearRequired
	}
	return nil
}

type BookFilter struct {
	Author string
	Page   int
	Limit  int
}

type PaginatedBooks struct {
	Data  []Book `json:"data"`
	Total int    `json:"total"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	Pages int    `json:"pages"`
}
