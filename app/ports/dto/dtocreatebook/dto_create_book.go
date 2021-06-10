package dtocreatebook

import "github.com/masihtehrani/books/app/entities/structs"

type Request struct {
	UserID string
	structs.Book
}

type Response struct {
	Success bool
}
