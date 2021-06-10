package mysql

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/masihtehrani/books/app/entities/structs"
)

func (m *Mysql) CreateBook(ctx context.Context, book structs.Book, userID string) error {
	book.ID = uuid.New().String()

	stmt, err := m.db.PrepareContext(ctx, "insert into books(id, author_id, title, description) values (?,?,?,?)")
	if err != nil {
		return fmt.Errorf("mysql.CreateBook >> %w", err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, book.ID, userID, book.Title, book.Description)
	if err != nil {
		return fmt.Errorf("mysql.CreateBook >> %w", err)
	}

	return nil
}
