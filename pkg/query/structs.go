package query

import (
	"database/sql"

	"github.com/masihtehrani/books/app/entities/structs"
)

type Where struct {
	Field     string
	Values    []string
	Operation string
}

type Query struct {
	Select       []string
	Body         string
	Query        string
	metaQuery    string
	GroupBy      string
	Operations   map[string]string
	wheres       []Where
	QueryFilters structs.Query
	meta         structs.Meta
	db           *sql.DB
	args         []interface{}
	maxPageSize  int
}
