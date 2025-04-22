package DatabaseService

import (
	"book-store-server/models"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

func (s *DBService) CreateBook(title, isbn, description, author string, updatedAt time.Time) (uint, error) {
	var id uint
	query := `
		INSERT INTO books (title, isbn, description, author, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	err := s.DB.QueryRow(query, title, isbn, description, author, updatedAt).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("error inserting book: %v", err)
	}

	return id, nil
}

func (s *DBService) UpdateBook(id uint, title, isbn, description, author string, updatedAt time.Time) error {
	query := `
		UPDATE books 
		SET title = $1, isbn = $2, description = $3, author = $4, updated_at = $5
		WHERE id = $6`

	_, err := s.DB.Exec(query, title, isbn, description, author, updatedAt, id)

	if err != nil {
		return fmt.Errorf("error updating book: %v", err)
	}

	return nil
}

func (s *DBService) GetBook(id uint) (*models.Book, error) {
	query := `
		SELECT id, title, isbn, description, author, created_at, updated_at 
		FROM books 
		WHERE id = $1`
	var book models.Book

	err := s.DB.QueryRow(query, id).Scan(&book.ID, &book.Title, &book.ISBN, &book.Description, &book.Author, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("book with ID %d not found", id)
		}
		return nil, fmt.Errorf("error fetching book: %v", err)
	}

	return &book, nil
}

func (s *DBService) DeleteBook(id uint) error {
	query := `DELETE FROM books WHERE id = $1`

	_, err := s.DB.Exec(query, id)

	if err != nil {
		return fmt.Errorf("error deleting book: %v", err)
	}

	return nil
}

func (s *DBService) GetBooks(limit, page int) ([]models.Book, int, error) {
	offset := (page - 1) * limit
	var books []models.Book

	query := `
		SELECT id, title, isbn, description, author, created_at, updated_at 
		FROM books 
		ORDER BY id DESC 
		LIMIT $1 OFFSET $2`
	rows, err := s.DB.Query(query, limit, offset)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Title, &book.ISBN, &book.Description, &book.Author, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, 0, err
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error processing rows: %v", err)
		return nil, 0, err
	}

	var totalBooks int
	countQuery := `SELECT COUNT(*) FROM books`
	err = s.DB.QueryRow(countQuery).Scan(&totalBooks)
	if err != nil {
		log.Printf("Error getting total book count: %v", err)
		return nil, 0, err
	}

	totalPages := (totalBooks + limit - 1) / limit

	return books, totalPages, nil
}
