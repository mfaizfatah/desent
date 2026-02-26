package domain

import "errors"

var (
	ErrBookNotFound   = errors.New("book not found")
	ErrTitleRequired  = errors.New("title is required")
	ErrAuthorRequired = errors.New("author is required")
	ErrYearRequired   = errors.New("year is required")
	ErrInvalidJSON    = errors.New("invalid JSON")
	ErrUnauthorized   = errors.New("invalid or expired token")
	ErrCredentials    = errors.New("username and password required")
)
