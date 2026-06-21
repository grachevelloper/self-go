package book

import (
	domainbook "book-service/internal/domain/book"
	"book-service/internal/domain/shared"
	usecasebook "book-service/internal/usecase/book"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type Service interface {
	Create(
		ctx context.Context,
		input usecasebook.CreateBookInput,
	) (*domainbook.Book, error)
	Update(
		ctx context.Context,
		input usecasebook.UpdateBookInput,
	) (*domainbook.Book, error)
	GetById(ctx context.Context, id string) (*domainbook.Book, error)
	GetAll(ctx context.Context) ([]*domainbook.Book, error)
	Delete(ctx context.Context, id string) error
}

var _ Service = (*usecasebook.UseCase)(nil)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
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
		var validationError *shared.ValidationError

		if errors.As(err, &validationError) {
			http.Error(w, "invalid book data", http.StatusBadRequest)
			return
		}
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

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var request UpdateBookRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalied request body", http.StatusBadRequest)
		return
	}

	input := usecasebook.UpdateBookInput{
		ID:          request.ID,
		Title:       request.Title,
		Author:      request.Author,
		Status:      request.Status,
		PublishedAt: request.PublishedAt,
	}

	updatedBook, err := h.service.Update(r.Context(), input)
	if err != nil {
		var validationError *shared.ValidationError

		if errors.As(err, &validationError) {
			http.Error(w, "invalid book data", http.StatusBadRequest)
			return
		}
		http.Error(w, "failed to create book", http.StatusInternalServerError)
		return
	}

	response := bookResponseMapper(updatedBook)
	w.WriteHeader(http.StatusAccepted)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		return
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.Create(w, r)
		case http.MethodPut:
			h.Update(w, r)
		}
	})
}
