package book

import (
	domainbook "book-service/internal/domain/book"
	"book-service/internal/domain/shared"
	usecasebook "book-service/internal/usecase/book"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
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
	logger  *slog.Logger
}

func NewHandler(service Service, logger *slog.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
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
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.logEncodeError(r, err)
		return
	}
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var request UpdateBookRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalied request body", http.StatusBadRequest)
		return
	}

	input := usecasebook.UpdateBookInput{
		ID:          id,
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
		h.logEncodeError(r, err)
		return
	}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {

	books, err := h.service.GetAll(r.Context())
	if err != nil {
		http.Error(w, "failed to create book", http.StatusInternalServerError)
		return
	}

	var response []BookResponse
	for _, book := range books {
		response = append(response, bookResponseMapper(book))
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.logEncodeError(r, err)
		return
	}
}

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	book, err := h.service.GetById(r.Context(), id)
	if err != nil {
		http.Error(w, "failed to get book by id", http.StatusInternalServerError)
		return
	}

	responseBook := bookResponseMapper(book)

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(responseBook); err != nil {
		h.logEncodeError(r, err)
		return
	}
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := h.service.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, "failed to create book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) logEncodeError(r *http.Request, err error) {
	if h.logger == nil {
		return
	}

	h.logger.Error("encode response failed",
		"error", err,
		"method", r.Method,
		"path", r.URL.Path,
	)
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.Create(w, r)
		case http.MethodGet:
			h.GetAll(w, r)
		}
	})

	mux.HandleFunc("/api/v1/books/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.Update(w, r)
		case http.MethodGet:
			h.GetById(w, r)
		case http.MethodDelete:
			h.Delete(w, r)
		}
	})

}
