package book

import (
	domainbook "book-service/internal/domain/book"
	usecasebook "book-service/internal/usecase/book"
	"context"
	"encoding/json"
	"net/http"
)

type Service interface {
	Create(
		ctx context.Context,
		input usecasebook.CreateBookInput,
	) (*domainbook.Book, error)
	Update(
		ctx context.Context,
		id int64,
		input usecasebook.UpdateBookInput,
	) (*domainbook.Book, error)
	GetById(ctx context.Context, id int64) (*domainbook.Book, error)
	GetAll(ctx context.Context) ([]*domainbook.Book, error)
	Delete(ctx context.Context, id int64) error
}

type Handler struct {
	service Service
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var request CreateBookRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	input := usecasebook.CreateBookInput{
		Title:       request.Title,
		Author:      request.Author,
		Status:      request.Status,
		PublishedAt: request.PublishedAt,
	}

	createdBook, err := h.service.Create(r.Context(), input)
	if err != nil {
		http.Error(w, "failed to create book", http.StatusInternalServerError)
		return
	}

	response := bookResponseMapper(createdBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		return
	}
}
