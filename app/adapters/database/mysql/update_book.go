package mysql

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/masihtehrani/books/app/entities/structs"
)

func (m *Mysql) UpdateBook(ctx context.Context, book structs.Book, userID, bookID string) error {
	book.ID = uuid.New().String()

	stmt, err := m.db.PrepareContext(ctx, "update books set "+
		"title = ?,description = ?,is_published = ? "+
		"where author_id = ? and id = ?")
	if err != nil {
		return fmt.Errorf("mysql.UpdateBook >> %w", err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, book.Title, book.Description, book.IsPublished, userID, bookID)
	if err != nil {
		return fmt.Errorf("mysql.UpdateBook >> %w", err)
	}

	return nil
}
