package repository

import (
	"book-service/internal/delivery/http/book"
	"book-service/internal/domain"
	"context"
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, book book.CreateBookInput) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO books (title, author, status)
		VALUES ($1, $2, $3)
	`, book.Title, book.Author, book.Status)

	return err
}

func (r *Repository) GetById(ctx context.Context, id int64) (*domain.Book, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT * 
		from books
		WHERE id = $1
	`, id)

	var book domain.Book

	if err := row.Scan(&book.ID, &book.Title, &book.Status, &book.Author); err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *Repository) GetAll(ctx context.Context) ([]*domain.Book, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT *
		from books
	`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []*domain.Book

	for rows.Next() {
		var book domain.Book

		if err := rows.Scan(&book.ID, &book.Title, &book.Status, &book.Author); err != nil {
			return nil, err
		}

		books = append(books, &book)

	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (r *Repository) Update(ctx context.Context, id int64, input book.UpdateBookInput) (book.UpdateBookOutput, error) {
	updatedBook := &domain.Book{
		ID: id,
	}

	if input.Title != nil {
		updatedBook.Title = *input.Title
	}

	if input.Author != nil {
		updatedBook.Author = *input.Author
	}
	if input.Status != nil {
		updatedBook.Status = *input.Status
	}

	row := r.db.QueryRowContext(ctx, `
		Update books		
		SET title = $1, author = $2, status = $3
		WHERE id = $4
		RETURNING id, title, status, author
	`, updatedBook.Title, updatedBook.Author, updatedBook.Status, id)

	if err := row.Scan(&updatedBook.ID, &updatedBook.Title, &updatedBook.Status, &updatedBook.Author); err != nil {
		return nil, err
	}

	return updatedBook, nil
}
