package book

import (
	httpShared "book-service/internal/delivery/http/shared"
	"book-service/internal/domain/book"
	domainbook "book-service/internal/domain/book"
	"book-service/internal/domain/shared"
	usecasebook "book-service/internal/usecase/book"
	"book-service/internal/usecase/shared/paginated"
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
	GetAll(ctx context.Context, input usecasebook.GetAllBooksInput) (*paginated.New[book.Book], error)
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
	page, err := httpShared.ParseIntQuery(r, "page")
	if err != nil {
		http.Error(w, "invalid page query parameter", http.StatusBadRequest)
		return
	}

	limit, err := httpShared.ParseIntQuery(r, "limit")
	if err != nil {
		http.Error(w, "invalid limit query parameter", http.StatusBadRequest)
		return
	}

	order, err := httpShared.ParseSortOrder(r.URL.Query().Get("order"))
	if err != nil {
		http.Error(w, "invalid order query parameter", http.StatusBadRequest)
		return
	}

	sortField, err := parseBookSortField(r.URL.Query().Get("sort_field"))
	if err != nil {
		http.Error(w, "invalid sort_field query parameter", http.StatusBadRequest)
		return
	}

	input := usecasebook.GetAllBooksInput{
		Page:      page,
		Limit:     limit,
		Order:     order,
		SortField: sortField,
	}

	paginatedBooks, err := h.service.GetAll(r.Context(), input)
	if err != nil {
		var validationError *shared.ValidationError
		if errors.As(err, &validationError) {
			http.Error(w, "invalid pagination data", http.StatusBadRequest)
			return
		}
		http.Error(w, "failed to get books", http.StatusInternalServerError)
		h.logEncodeError(r, err)
		return
	}

	var books []BookResponse
	for _, book := range paginatedBooks.Items {
		books = append(books, bookResponseMapper(&book))
	}

	response := httpShared.PaginatedResponseMapper(paginatedBooks, books)

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

func parseBookSortField(order string) (usecasebook.BookSortField, error) {
	if order == "" {
		return usecasebook.CreatedAt, nil
	}
	sortField := usecasebook.BookSortField(order)

	switch sortField {
	case usecasebook.Title, usecasebook.CreatedAt:
		return sortField, nil
	default:
		return "", &shared.ValidationError{
			Field: "sort_field",
			Code:  "invalid",
		}
	}
}
