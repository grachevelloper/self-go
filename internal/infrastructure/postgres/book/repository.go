package repository

import (
	"book-service/internal/domain/book"
	domain "book-service/internal/domain/book"
	"book-service/internal/domain/shared/paginated"
	usecasebook "book-service/internal/usecase/book"
	"context"
	"database/sql"
	"errors"
	"time"
)

const bookColumns = `
	id,
	title,
	author,
	status,
	published_at,
	created_at,
	updated_at
`

type Repository struct {
	db *sql.DB
}

// compile-time проверка, что PostgreSQL-репозиторий реализует интерфейс use case
var _ usecasebook.Repository = (*Repository)(nil)

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, entity *domain.Book) (*domain.Book, error) {
	row := r.db.QueryRowContext(ctx, `
		INSERT INTO books (
			id,
			title,
			author,
			status,
			published_at,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING `+bookColumns+`
	`,
		entity.ID(),
		entity.Title(),
		entity.Author(),
		entity.Status(),
		entity.PublishedAt(),
		entity.CreatedAt(),
		entity.UpdatedAt(),
	)

	return scanBook(row)
}

func (r *Repository) GetById(ctx context.Context, id string) (*domain.Book, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT `+bookColumns+`
		FROM books
		WHERE id = $1
	`, id)

	return scanBook(row)
}

func (r *Repository) GetAll(ctx context.Context, pE paginated.PaginationParams) (*paginated.PaginatedEntity[book.Book], error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT `+bookColumns+`
		FROM books
		Limit $1
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []domain.Book

	for rows.Next() {
		book, err := scanBook(rows)
		if err != nil {
			return nil, err
		}

		books = append(books, *book)

	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	var total int
	row := r.db.QueryRowContext(ctx, `
		SELECT Count(*)
		FROM books
	`)

	if err := row.Scan(
		&total,
	); err != nil {
		return nil, err
	}

	hasNext := total > len(books)

	return &paginated.PaginatedEntity[domain.Book]{
		Items:   books,
		HasNext: hasNext,
		Total:   total,
		Page:    pE.Page,
		Limit:   pE.Limit,
	}, nil
}

func (r *Repository) Update(
	ctx context.Context,
	id string,
	params domain.UpdateBookParams,
	expectedUpdatedAt *time.Time,
) (*domain.Book, error) {
	var status any
	if params.Status != nil {
		status = string(*params.Status)
	}

	row := r.db.QueryRowContext(ctx, `
		UPDATE books
		SET
			title = COALESCE($1, title),
			author = COALESCE($2, author),
			status = COALESCE($3, status),
			published_at = COALESCE($4, published_at),
			updated_at = $5
		WHERE id = $6
			AND updated_at IS NOT DISTINCT FROM $7
		RETURNING `+bookColumns+`
	`,
		params.Title,
		params.Author,
		status,
		params.PublishedAt,
		params.UpdatedAt,
		id,
		expectedUpdatedAt,
	)

	updated, err := scanBook(row)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, usecasebook.ErrConcurrentUpdate
	}
	return updated, err
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, `
		DELETE FROM books
		WHERE id=$1
	`, id)
	if err != nil {
		return err
	}

	return nil
}

type scanner interface {
	Scan(dest ...any) error
}

func scanBook(s scanner) (*domain.Book, error) {
	var (
		id          string
		title       string
		author      string
		status      domain.BookStatus
		publishedAt time.Time
		createdAt   time.Time
		updatedAt   *time.Time
	)

	err := s.Scan(
		&id,
		&title,
		&author,
		&status,
		&publishedAt,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}

	return domain.RestoreBook(domain.NewBookParams{
		ID:          id,
		Title:       title,
		Author:      author,
		Status:      status,
		PublishedAt: publishedAt,
		CreatedAt:   createdAt,
	}, updatedAt)
}
