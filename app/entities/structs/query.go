package structs

type Meta struct {
	PageNumber int `json:"current_page"`
	PageSize   int `json:"per_page"`
	LastPage   int `json:"last_page"`
	Total      int `json:"total"`
}

type Query struct {
	Filter map[string]map[string][]string // [field_name][]values
	Sort   map[string][]string            // [asc | des][]field_name
	Page
}

type Page struct {
	Number int
	Size   int
}
