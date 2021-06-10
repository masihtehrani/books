package dtogetbooks

import "github.com/masihtehrani/books/app/entities/structs"

type Request struct {
	structs.Query
}

type Response struct {
	Books []structs.Book `json:"data"`
	Meta  structs.Meta   `json:"meta"`
}
