package mysql

import "fmt"

func (m *Mysql) Close() error {
	if err := m.db.Close(); err != nil {
		return fmt.Errorf("Close() >> %w", err)
	}

	return nil
}
