package response

import "book-store-server/models"

type CreateBookResponse struct {
	Book *models.Book `json:"book"`
}

type GetBooksResponse struct {
	Books []models.Book `json:"books"`
}

type GetBookResponse struct {
	Book *models.Book `json:"book"`
}
