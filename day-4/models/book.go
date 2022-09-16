package models

type Book struct {
	ID        string `json:"id"`
	Title     string `json:"title" form:"title" validate:"required"`
	ISBN      string `json:"isbn" form:"isbn"`
	Writer    string `json:"writer" form:"writer" validate:"required"`
	CreatedBy int    `json:"created_by"`
}
