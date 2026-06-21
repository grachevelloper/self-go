package book

import (
	usecasebook "book-service/internal/usecase/book"
	"book-service/internal/usecase/book/mocks"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestCreateReturnsBadRequestForValidationError(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := mocks.NewMockRepository(ctrl)
	service := usecasebook.NewUseCase(repository, func() string { return "book-id" })
	handler := &Handler{service: service}
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
}
