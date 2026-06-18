package book

type BookStatus string

const (
	BookStatusDraft     BookStatus = "reading"
	BookStatusPublished BookStatus = "in_wishlist"
	BookStatusArchived  BookStatus = "finished"
)
