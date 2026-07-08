package book

import (
	httpShared "book-service/internal/delivery/http/shared"
	domainbook "book-service/internal/domain/book"
	usecasebook "book-service/internal/usecase/book"
	"book-service/internal/usecase/book/mocks"
	"book-service/internal/usecase/shared/order"
	"book-service/internal/usecase/shared/paginated"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"go.uber.org/mock/gomock"
)

func TestCreate(t *testing.T) {
	t.Run("Bad request return validation error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repository := mocks.NewMockRepository(ctrl)
		service := usecasebook.NewUseCase(repository, func() string { return "book-id" })
		handler := NewHandler(service, slog.New(slog.NewTextHandler(io.Discard, nil)))
		request := httptest.NewRequest(http.MethodPost, "/books", bytes.NewBufferString(`{
			"title":"",
			"author":"George Orwell",
			"status":"reading",
			"published_at":"1949-06-08T00:00:00Z"
		}`))
		response := httptest.NewRecorder()

		handler.Create(response, request)

		if response.Code != http.StatusBadRequest {
			t.Fatalf("Create() status = %d, want %d", response.Code, http.StatusBadRequest)
		}
	})

	t.Run("return book for valid data", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repository := mocks.NewMockRepository(ctrl)
		service := usecasebook.NewUseCase(repository, func() string { return "book-id" })
		handler := NewHandler(service, slog.New(slog.NewTextHandler(io.Discard, nil)))

		request := httptest.NewRequest(http.MethodPost, "/books", bytes.NewBufferString(`{
			"title":"Some okey title",
			"author":"George Orwell",
			"status":"reading",
			"published_at":"1949-06-08T00:00:00Z"
		}`))
		response := httptest.NewRecorder()

		repository.EXPECT().
			Create(gomock.Any(), gomock.Any()).
			DoAndReturn(func(_ context.Context, entity *domainbook.Book) (*domainbook.Book, error) {
				return entity, nil
			})

		handler.Create(response, request)

		if response.Code != http.StatusCreated {
			t.Fatalf("Create() status = %d, want %d", response.Code, http.StatusCreated)
		}
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Bad request return validation error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repository := mocks.NewMockRepository(ctrl)
		service := usecasebook.NewUseCase(repository, func() string { return "book-id" })
		handler := NewHandler(service, slog.New(slog.NewTextHandler(io.Discard, nil)))
		request := httptest.NewRequest(http.MethodPost, "/books/{id}", bytes.NewBufferString(`{
			"title":"",
			"author":"George Orwell",
			"status":"reading",
			"published_at":"1949-06-08T00:00:00Z"
		}`))
		response := httptest.NewRecorder()

		handler.Update(response, request)

		if response.Code != http.StatusBadRequest {
			t.Fatalf("Update() status = %d, want %d", response.Code, http.StatusBadRequest)
		}
	})

	t.Run("return book for valid data", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repository := mocks.NewMockRepository(ctrl)
		service := usecasebook.NewUseCase(repository, func() string { return "book-id" })
		handler := NewHandler(service, slog.New(slog.NewTextHandler(io.Discard, nil)))

		request := httptest.NewRequest(http.MethodPost, "/books", bytes.NewBufferString(`{
			"title":"Some okey title",
			"author":"George Orwell",
			"status":"reading",
			"published_at":"1949-06-08T00:00:00Z"
		}`))
		request.SetPathValue("id", "book-id")
		response := httptest.NewRecorder()

		existing, err := domainbook.RestoreBook(domainbook.NewBookParams{
			ID:          "book-id",
			Title:       "Old title",
			Author:      "George Orwell",
			Status:      domainbook.BookStatus("reading"),
			PublishedAt: time.Date(1949, time.June, 8, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC),
		}, nil)
		if err != nil {
			t.Fatalf("RestoreBook() error = %v", err)
		}

		repository.EXPECT().
			GetById(gomock.Any(), "book-id").
			Return(existing, nil)

		repository.EXPECT().
			Update(gomock.Any(), "book-id", gomock.Any(), gomock.Any()).
			DoAndReturn(func(_ context.Context, _ string, params domainbook.UpdateBookParams, _ *time.Time) (*domainbook.Book, error) {
				return existing.Updated(params)
			})

		handler.Update(response, request)

		if response.Code != http.StatusAccepted {
			t.Fatalf("Update() status = %d, want %d", response.Code, http.StatusAccepted)
		}
	})
}

func TestGetAll(t *testing.T) {
	t.Run("valid data", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repository := mocks.NewMockRepository(ctrl)
		service := usecasebook.NewUseCase(repository, func() string { return "unused" })
		handler := NewHandler(service, slog.New(slog.NewTextHandler(io.Discard, nil)))

		request := httptest.NewRequest(http.MethodGet, "/books?page=2&limit=3&order=asc&sort_field=title", nil)
		response := httptest.NewRecorder()

		createdAt := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
		book, err := domainbook.RestoreBook(domainbook.NewBookParams{
			ID:          "book-id",
			Title:       "1984",
			Author:      "George Orwell",
			Status:      domainbook.BookStatus("reading"),
			PublishedAt: time.Date(1949, time.June, 8, 0, 0, 0, 0, time.UTC),
			CreatedAt:   createdAt,
		}, nil)
		if err != nil {
			t.Fatalf("RestoreBook() error = %v", err)
		}

		repository.EXPECT().
			GetAll(gomock.Any(), paginated.PaginationParams[usecasebook.BookSortField]{
				Page:      2,
				Limit:     3,
				Order:     order.Asc,
				SortField: usecasebook.Title,
			}).
			Return(&paginated.New[domainbook.Book]{
				Items:   []domainbook.Book{*book},
				Page:    2,
				Limit:   3,
				Total:   4,
				HasNext: true,
			}, nil)

		handler.GetAll(response, request)

		if response.Code != http.StatusOK {
			t.Fatalf("GetAll() status = %d, want %d", response.Code, http.StatusOK)
		}

		var got httpShared.PaginatedResponse[BookResponse]
		if err := json.NewDecoder(response.Body).Decode(&got); err != nil {
			t.Fatalf("Decode() error = %v", err)
		}
		if got.Page != 2 || got.Limit != 3 || got.Total != 4 || !got.HasNext {
			t.Fatalf("GetAll() response = %+v, want page=2 limit=3 total=4 has_next=true", got)
		}
		if len(got.Items) != 1 || got.Items[0].ID != "book-id" {
			t.Fatalf("GetAll() items = %+v, want one book-id item", got.Items)
		}
	})
}
