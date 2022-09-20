package entity

type Book struct {
	ID        string `json:"id"`
	Title     string `json:"title" form:"title" validate:"required"`
	ISBN      string `json:"isbn" form:"isbn"`
	Writer    string `json:"writer" form:"writer" validate:"required"`
	CreatedBy int    `json:"created_by"`
}

type BookRepository interface {
	CreateBook(book Book) (interface{}, error)
	GetBooks() (interface{}, error)
	GetBookByID(id string) (interface{}, error)
	UpdateBook(book Book, id string) (interface{}, error)
	DeleteBook(id string) (interface{}, error)
}
