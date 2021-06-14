package mysql

import (
	"context"
	"fmt"
)

func (m *Mysql) DeleteBook(ctx context.Context, userID, bookID string) error {
	stmt, err := m.db.PrepareContext(ctx, "delete from books where author_id = ? and id = ?")
	if err != nil {
		return fmt.Errorf("mysql.DeleteBook >> %w", err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, userID, bookID)
	if err != nil {
		return fmt.Errorf("mysql.DeleteBook >> %w", err)
	}

	return nil
}
