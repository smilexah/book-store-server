package services

import (
	"book-store-server/models"
	"book-store-server/models/request"
	"book-store-server/models/response"
	"fmt"
	"net/http"
	"time"
)

func (s *AppService) CreateBook(req *request.CreateBookRequest) (*response.Response, int) {
	newBookId, err := s.DBService.CreateBook(req.Title, req.ISBN, req.Description, req.Author, time.Now())
	if err != nil {
		return &response.Response{
			Error: err.Error(),
		}, http.StatusInternalServerError
	}

	createBookRes := response.CreateBookResponse{
		Book: &models.Book{
			ID:          newBookId,
			Title:       req.Title,
			ISBN:        req.ISBN,
			Description: req.Description,
			Author:      req.Author,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	return &response.Response{
		Result: createBookRes,
	}, http.StatusCreated
}

func (s *AppService) GetBook(id uint) (*response.Response, int) {
	book, err := s.DBService.GetBook(id)
	if err != nil {
		return &response.Response{
			Error: err.Error(),
		}, http.StatusNotFound
	}

	bookRes := response.GetBookResponse{
		Book: book,
	}

	return &response.Response{
		Result: bookRes,
	}, http.StatusOK
}

func (s *AppService) DeleteBook(id uint) (*response.Response, int) {
	err := s.DBService.DeleteBook(id)
	if err != nil {
		return &response.Response{
			Error: err.Error(),
		}, http.StatusInternalServerError
	}

	return &response.Response{
		Result: "Book deleted successfully",
	}, http.StatusNoContent
}

func (s *AppService) UpdateBook(req *request.UpdateBookRequest) (*response.Response, int) {
	err := s.DBService.UpdateBook(req.ID, req.Title, req.ISBN, req.Description, req.Author, time.Now())

	if err != nil {
		return &response.Response{
			Error: err.Error(),
		}, http.StatusInternalServerError
	}

	return &response.Response{
		Result: "Book updated successfully",
	}, http.StatusOK
}

func (s *AppService) GetBooks(limit, page int) (*response.Response, int) {
	books, totalPages, err := s.DBService.GetBooks(limit, page)
	if err != nil {
		return &response.Response{
			Error: err.Error(),
		}, http.StatusInternalServerError
	}

	fmt.Println(books)

	getBooksRes := &response.GetBooksResponse{
		Books: books,
	}

	paginationRes := response.PaginationResponse{
		Result:     getBooksRes,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}

	return &response.Response{
		Result: paginationRes,
	}, http.StatusOK
}
