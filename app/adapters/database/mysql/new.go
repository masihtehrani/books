package mysql

import (

	// import mysql.
	"context"
	"database/sql"
	"fmt"
	"os"

	// mysql driver.
	_ "github.com/go-sql-driver/mysql"
	"github.com/masihtehrani/books/app/entities/interfaces"
	"github.com/masihtehrani/books/pkg/logger"
)

type Mysql struct {
	db *sql.DB
}

type RawConfigMap map[string][]byte

func New(ctx context.Context, logger *logger.Logger) (interfaces.IDatabase, error) {
	dbName := os.Getenv("MYSQL_DBNAME")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASS")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, dbName))
	if err != nil {
		return nil, fmt.Errorf("mysql.New >> %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("mysql.New >> %w", err)
	}

	return &Mysql{db: db}, nil
}
