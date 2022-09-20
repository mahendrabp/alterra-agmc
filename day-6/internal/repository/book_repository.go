package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mahendrabp/alterra-agmc/day-6/internal/entity"
	"gorm.io/gorm"
)

type book struct {
}

func NewBookRepository(db *gorm.DB) entity.BookRepository {
	return &book{}
}

var books = map[string]entity.Book{
	"3123-3123": {
		ID:        "3123-3123",
		Title:     "book 1",
		ISBN:      "123-321",
		Writer:    "writer-1",
		CreatedBy: 0,
	},
	"3123-3124": {
		ID:        "3123-3123",
		Title:     "book 2",
		ISBN:      "123-322",
		Writer:    "writer-2",
		CreatedBy: 0,
	},
}

func (b *book) CreateBook(book entity.Book) (interface{}, error) {
	id := uuid.New().String()

	book.ID = id
	books[id] = book

	return books[id], nil

}

func (b *book) GetBooks() (interface{}, error) {
	var bookSlice []entity.Book

	for _, v := range books {
		bookSlice = append(bookSlice, v)
	}

	return bookSlice, nil
}

func (b *book) GetBookByID(id string) (interface{}, error) {
	bookSlice := books[id]

	if bookSlice.ID == "" {
		return nil, errors.New("data not found")
	}

	return bookSlice, nil
}

func (b *book) UpdateBook(book entity.Book, id string) (interface{}, error) {

	bookSlice := books[id]

	if bookSlice.ID == "" {
		return nil, errors.New("data not found")
	}

	book.ID = id
	books[id] = book

	return books[id], nil
}

func (b *book) DeleteBook(id string) (interface{}, error) {
	delete(books, id)
	return nil, nil
}
