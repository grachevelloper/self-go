package book_test

import (
	domainbook "book-service/internal/domain/book"
	usecasebook "book-service/internal/usecase/book"
	mocks "book-service/internal/usecase/book/mocks"
	"context"
	"testing"
	"time"

	"go.uber.org/mock/gomock"
)

var mockedBookInput = usecasebook.CreateBookInput{
	Title:       "1984",
	Author:      "Оруэл",
	PublishedAt: time.Date(1947, time.June, 18, 0, 0, 0, 0, time.UTC),
	Status:      "reading",
}

func TestCreateUseCase(t *testing.T) {
	t.Run("valid input data", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repository := mocks.NewMockRepository(ctrl)
		useCase := usecasebook.NewUseCase(repository, func() string { return "book-id" })
		ctx := t.Context()

		repository.EXPECT().
			Create(ctx, gomock.Any()).
			DoAndReturn(func(_ context.Context, entity *domainbook.Book) (*domainbook.Book, error) {
				if entity.ID() != "book-id" {
					t.Errorf("Create() ID = %q, want %q", entity.ID(), "book-id")
				}
				if entity.Title() != mockedBookInput.Title || entity.Author() != mockedBookInput.Author {
					t.Errorf("Create() entity = %#v, want input fields", entity)
				}
				if entity.CreatedAt().IsZero() {
					t.Error("Create() did not set CreatedAt")
				}
				return entity, nil
			})

		got, err := useCase.Create(ctx, mockedBookInput)
		if err != nil {
			t.Fatalf("Create() error = %v", err)
		}
		if got.ID() != "book-id" {
			t.Errorf("Create() ID = %q, want %q", got.ID(), "book-id")
		}
	})
}

func TestUpdateUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := mocks.NewMockRepository(ctrl)
	useCase := usecasebook.NewUseCase(repository, func() string { return "unused" })
	ctx := t.Context()

	createdAt := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	previousUpdatedAt := time.Date(2025, time.February, 1, 0, 0, 0, 0, time.UTC)

	existing, err := domainbook.RestoreBook(domainbook.NewBookParams{
		ID:          "book-id",
		Title:       "Old title",
		Author:      "Author",
		Status:      domainbook.BookStatus("reading"),
		PublishedAt: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		CreatedAt:   createdAt,
	}, &previousUpdatedAt)

	if err != nil {
		t.Fatalf("NewBook() error = %v", err)
	}

	newTitle := "New title"
	input := usecasebook.UpdateBookInput{ID: existing.ID(), Title: &newTitle}

	repository.EXPECT().GetById(ctx, existing.ID()).Return(existing, nil)
	repository.EXPECT().Update(ctx, existing.ID(), gomock.Any(), &previousUpdatedAt).DoAndReturn(
		func(_ context.Context, id string, params domainbook.UpdateBookParams, expectedUpdatedAt *time.Time) (*domainbook.Book, error) {
			if id != existing.ID() {
				t.Errorf("Update() id = %q, want %q", id, existing.ID())
			}
			if expectedUpdatedAt == nil || !expectedUpdatedAt.Equal(previousUpdatedAt) {
				t.Errorf("Update() expectedUpdatedAt = %v, want %v", expectedUpdatedAt, previousUpdatedAt)
			}
			if params.Title == nil || *params.Title != newTitle {
				t.Errorf("Update() title param = %v, want %q", params.Title, newTitle)
			}
			if params.Author != nil || params.Status != nil || params.PublishedAt != nil {
				t.Error("Update() set params absent from input")
			}
			updated, err := existing.Updated(params)
			if err != nil {
				t.Fatalf("Updated() error = %v", err)
			}
			if updated == existing {
				t.Error("Update() mutated the repository entity instead of creating a copy")
			}
			if updated.CreatedAt() != createdAt {
				t.Error("Update() changed CreatedAt")
			}
			if updated.UpdatedAt() == nil {
				t.Fatal("Update() did not set UpdatedAt")
			}
			return updated, nil
		},
	)

	got, err := useCase.Update(ctx, input)
	if err != nil {
		t.Fatalf("Update() error = %v", err)
	}
	if got.Title() != newTitle {
		t.Errorf("Update() title = %q, want %q", got.Title(), newTitle)
	}
}
