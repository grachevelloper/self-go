package book

import (
	domainbook "book-service/internal/domain/book"
	usecasebook "book-service/internal/usecase/book"
	"book-service/internal/usecase/book/mocks"
	"bytes"
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

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
		response := httptest.NewRecorder()

		repository.EXPECT().
			Update(gomock.Any(), gomock.Any(), gomock.Any()).
			DoAndReturn(func(_ context.Context, entity *domainbook.Book) (*domainbook.Book, error) {
				return entity, nil
			})

		handler.Create(response, request)

		if response.Code != http.StatusCreated {
			t.Fatalf("Create() status = %d, want %d", response.Code, http.StatusCreated)
		}
	})
}
