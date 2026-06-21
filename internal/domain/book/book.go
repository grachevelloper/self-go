package book

import "time"

type Book struct {
	id          string
	title       string
	author      string
	status      BookStatus
	publishedAt time.Time
	createdAt   time.Time
	updatedAt   *time.Time
}

type NewBookParams struct {
	ID          string
	Title       string
	Author      string
	Status      BookStatus
	PublishedAt time.Time
	CreatedAt   time.Time
}

type UpdateBookParams struct {
	Title       *string
	Author      *string
	Status      *BookStatus
	PublishedAt *time.Time
	UpdatedAt   time.Time
}

func NewBook(params NewBookParams) (*Book, error) {
	if err := validateBook(params); err != nil {
		return nil, err
	}

	return &Book{
		id:          params.ID,
		title:       params.Title,
		author:      params.Author,
		status:      params.Status,
		publishedAt: params.PublishedAt,
		createdAt:   params.CreatedAt,
	}, nil
}

func RestoreBook(params NewBookParams, updatedAt *time.Time) (*Book, error) {
	entity, err := NewBook(params)
	if err != nil {
		return nil, err
	}
	if updatedAt != nil {
		value := *updatedAt
		entity.updatedAt = &value
	}
	return entity, nil
}

func (b *Book) Updated(params UpdateBookParams) (*Book, error) {
	updated := *b

	if params.Title != nil {
		if err := ValidateTitle(*params.Title); err != nil {
			return nil, err
		}
		updated.title = *params.Title
	}
	if params.Author != nil {
		if err := ValidateAuthor(*params.Author); err != nil {
			return nil, err
		}
		updated.author = *params.Author
	}
	if params.Status != nil {
		if err := ValidateStatus(string(*params.Status)); err != nil {
			return nil, err
		}
		updated.status = *params.Status
	}
	if params.PublishedAt != nil {
		if err := ValidatePublishedAt(*params.PublishedAt); err != nil {
			return nil, err
		}
		updated.publishedAt = *params.PublishedAt
	}

	updated.updatedAt = &params.UpdatedAt
	return &updated, nil
}

func validateBook(params NewBookParams) error {
	if err := ValidateID(params.ID); err != nil {
		return err
	}
	if err := ValidateTitle(params.Title); err != nil {
		return err
	}
	if err := ValidateAuthor(params.Author); err != nil {
		return err
	}
	if err := ValidateStatus(string(params.Status)); err != nil {
		return err
	}
	return ValidatePublishedAt(params.PublishedAt)
}

func (b *Book) ID() string             { return b.id }
func (b *Book) Title() string          { return b.title }
func (b *Book) Author() string         { return b.author }
func (b *Book) Status() BookStatus     { return b.status }
func (b *Book) PublishedAt() time.Time { return b.publishedAt }
func (b *Book) CreatedAt() time.Time   { return b.createdAt }

func (b *Book) UpdatedAt() *time.Time {
	if b.updatedAt == nil {
		return nil
	}
	value := *b.updatedAt
	return &value
}
