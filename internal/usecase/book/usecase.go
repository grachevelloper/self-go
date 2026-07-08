package book

import (
	"book-service/internal/domain/book"
	"book-service/internal/usecase/shared/paginated"
	"book-service/internal/usecase/shared/uuid"
	"context"
	"time"
)

type UseCase struct {
	repo        Repository
	generatorID uuid.IDGenerator
}

func NewUseCase(repo Repository, generateID uuid.IDGenerator) *UseCase {
	return &UseCase{
		repo:        repo,
		generatorID: generateID,
	}
}

func (u *UseCase) Create(ctx context.Context, input CreateBookInput) (*book.Book, error) {

	book, err := input.toBook(u.generatorID)

	if err != nil {
		return nil, err
	}

	return u.repo.Create(ctx, book)
}
func (u *UseCase) GetById(ctx context.Context, id string) (*book.Book, error) {
	return u.repo.GetById(ctx, id)
}

func (u *UseCase) GetAll(ctx context.Context, input GetAllBooksInput) (*paginated.New[book.Book], error) {

	paginatedParams := paginated.PaginationParams[BookSortField]{
		Page:      input.Page,
		Limit:     input.Limit,
		Order:     input.Order,
		SortField: input.SortField,
	}

	if err := paginated.ValidatePaginatedEntity(paginatedParams); err != nil {
		return nil, err
	}

	return u.repo.GetAll(ctx, paginatedParams)
}

func (u *UseCase) Update(ctx context.Context, input UpdateBookInput) (*book.Book, error) {
	if err := book.ValidateID(input.ID); err != nil {
		return nil, err
	}

	existing, err := u.repo.GetById(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	params := input.toUpdateBookParams()

	updated, err := existing.Updated(params)
	if err != nil {
		return nil, err
	}

	return u.repo.Update(ctx, updated.ID(), params, existing.UpdatedAt())
}

func (u *UseCase) Delete(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}

func (input CreateBookInput) toBook(generatorID uuid.IDGenerator) (*book.Book, error) {
	return book.NewBook(book.NewBookParams{
		ID:          generatorID(),
		Title:       input.Title,
		Author:      input.Author,
		Status:      book.BookStatus(input.Status),
		PublishedAt: input.PublishedAt,
		CreatedAt:   time.Now(),
	})
}

func (input UpdateBookInput) toUpdateBookParams() book.UpdateBookParams {
	var status *book.BookStatus
	if input.Status != nil {
		value := book.BookStatus(*input.Status)
		status = &value
	}

	return book.UpdateBookParams{
		Title:       input.Title,
		Author:      input.Author,
		Status:      status,
		PublishedAt: input.PublishedAt,
		UpdatedAt:   time.Now(),
	}
}
