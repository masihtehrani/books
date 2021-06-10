package mysql

import (
	"context"
	"fmt"
)

func (m *Mysql) SignIn(ctx context.Context, username, password string) (string, error) {
	stmt, err := m.db.PrepareContext(ctx, "select id from authors where username = ? and password = ?")
	if err != nil {
		return "", fmt.Errorf("mysql.SignIn >> %w", err)
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, username, password)
	if row.Err() != nil {
		return "", fmt.Errorf("mysql.SignIn >> %w", err)
	}

	var id string

	err = row.Scan(&id)
	if err != nil {
		return "", fmt.Errorf("mysql.SignIn >> %w", err)
	}

	return id, nil
}
