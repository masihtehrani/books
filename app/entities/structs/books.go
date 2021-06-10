package structs

import "time"

type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsPublished bool   `json:"is_published"`
	Author
	CreatedAt time.Time `json:"created_at"`
}

type Author struct {
	FullName  string `json:"author_full_name"`
	Pseudonym string `json:"author_pseudonym"`
	Username  string `json:"-"`
	Password  string `json:"-"`
}
