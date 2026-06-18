package book

import "errors"

var (
	ErrAlreadyArchived = errors.New("book is already archived")
	ErrCannotPublish   = errors.New("book cannot be published")
)
