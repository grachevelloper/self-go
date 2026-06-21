package main

import (
	"book-service/internal/delivery"
	"book-service/internal/delivery/http/book"
	postgresbook "book-service/internal/infrastructure/postgres/book"
	"book-service/internal/infrastructure/uuid"
	usecasebook "book-service/internal/usecase/book"
	"database/sql"
	"net/http"
)

func main() {
	repository := postgresbook.NewRepository(&sql.DB{})
	bookHandler := book.NewHandler(usecasebook.NewUseCase(repository, uuid.GenerateUUID))

	router := delivery.NewRouter(
		bookHandler,
	)

	http.ListenAndServe(":8080", router)
}
