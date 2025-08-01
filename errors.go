package feed

import "errors"

// Common errors
var (
	ErrMissingTitle       = errors.New("feed title is required")
	ErrMissingDescription = errors.New("feed description is required")
	ErrMissingLink        = errors.New("feed link is required")
	ErrMissingItemTitle   = errors.New("item title is required")
	ErrMissingItemLink    = errors.New("item link is required")
	ErrInvalidURL         = errors.New("invalid URL format")
	ErrInvalidDate        = errors.New("invalid date format")
	ErrEmptyFeed          = errors.New("feed contains no items")
)
