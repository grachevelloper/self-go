package book

import "errors"

var ErrConcurrentUpdate = errors.New("book was concurrently updated")
