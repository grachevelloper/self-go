package repository

import (
	domain "book-service/internal/domain/book"
	usecase "book-service/internal/usecase/book"
	"context"
	"database/sql"
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

var _ usecase.Repository = (*Repository)(nil)

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, input usecase.CreateBookInput) (*domain.Book, error) {
	row := r.db.QueryRowContext(ctx, `
		INSERT INTO books (title, author, status, published_at)
		VALUES ($1, $2, $3, $4)
		RETURNING `+bookColumns+`
	`, input.Title, input.Author, input.Status, input.PublishedAt)

	return scanBook(row)
}

func (r *Repository) GetById(ctx context.Context, id int64) (*domain.Book, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT `+bookColumns+`
		FROM books
		WHERE id = $1
	`, id)

	return scanBook(row)
}

func (r *Repository) GetAll(ctx context.Context) ([]*domain.Book, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT `+bookColumns+`
		FROM books
	`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []*domain.Book

	for rows.Next() {
		book, err := scanBook(rows)
		if err != nil {
			return nil, err
		}

		books = append(books, book)

	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (r *Repository) Update(ctx context.Context, id int64, input usecase.UpdateBookInput) (*domain.Book, error) {
	row := r.db.QueryRowContext(ctx, `
		UPDATE books
		SET
			title = COALESCE($1, title),
			author = COALESCE($2, author),
			status = COALESCE($3, status),
			published_at = COALESCE($4, published_at),
			updated_at = now()
		WHERE id = $5
		RETURNING `+bookColumns+`
	`, input.Title, input.Author, input.Status, input.PublishedAt, id)

	return scanBook(row)
}

func (r *Repository) Delete(ctx context.Context, id int64) error {
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
	var book domain.Book

	err := s.Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.Status,
		&book.PublishedAt,
		&book.CreatedAt,
		&book.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &book, nil
}
