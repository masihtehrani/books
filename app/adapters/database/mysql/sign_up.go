package mysql

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/masihtehrani/books/app/entities/structs"
)

func (m *Mysql) SignUp(ctx context.Context, author structs.Author) error {
	userID := uuid.New().String()

	stmt, err := m.db.PrepareContext(ctx, ""+
		"insert into authors(id, full_name, pseudonym, "+
		"username, password) "+
		"values(?,?,?,?,?)")
	if err != nil {
		return fmt.Errorf("mysql.SignUp >> %w", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, userID, author.FullName, author.Pseudonym, author.Username, author.Password)
	if err != nil {
		return fmt.Errorf("mysql.SignUp >> %w", err)
	}

	return nil
}
