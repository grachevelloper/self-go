package domain

type BookStatus string

const (
	BookStatusDraft     BookStatus = "reading"
	BookStatusPublished BookStatus = "in_wishlist"
	BookStatusArchived  BookStatus = "finished"
)

func (s BookStatus) Valid() bool {
	switch s {
	case BookStatusDraft, BookStatusPublished, BookStatusArchived:
		return true
	default:
		return false
	}
}
