package testdata_test

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	mysql2 "github.com/masihtehrani/books/app/adapters/database/mysql"
	"github.com/masihtehrani/books/app/ports/dto/dtoversion"
	http2 "github.com/masihtehrani/books/app/ports/transport/http"
	"github.com/masihtehrani/books/app/usecases"
	"github.com/masihtehrani/books/pkg/logger"
	"github.com/ory/dockertest/v3"
)

var (
	dbName = "books"
	dbPass = "secret"
	dbUser = "root"
	dbHost = "127.0.0.1"
	dbPort = "3306"

	ip   = "127.0.0.1"
	port = "6000"

	token = ""
)

type BooksTests struct{ Test *testing.T }

func TestRunner(t *testing.T) {

	t.Run("Api Test", func(t *testing.T) {
		test := BooksTests{Test: t}
		test.TestSignUp()
		test.TestSignIn()
		test.TestCreateBook()
	})
}

func TestMain(m *testing.M) {
	var db *sql.DB

	os.Setenv("HTTP_IP", ip)
	os.Setenv("HTTP_PORT", port)
	os.Setenv("MYSQL_DBNAME", dbName)
	os.Setenv("MYSQL_HOST", dbHost)
	os.Setenv("MYSQL_PORT", dbPort)
	os.Setenv("MYSQL_USER", dbUser)
	os.Setenv("MYSQL_PASS", dbPass)
	os.Setenv("MYSQL_SSLMODE", "false")
	os.Setenv("JWT_SECRET_KEY", "699f8ece-3c5f-4e3b-b877-090870428635")
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("mariadb", "10.6", []string{"MYSQL_ROOT_PASSWORD=" + dbPass, "MYSQL_DATABASE=" + dbName})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	dbPort = resource.GetPort("3306/tcp")
	os.Setenv("MYSQL_PORT", dbPort)

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error
		db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	mm, err := migrate.New(
		"file://../migrations",
		fmt.Sprintf("mysql://%s:%s@(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName))
	if err != nil {
		log.Fatalf("migration: %s", err)
	}
	mm.Steps(2)

	res, _ := db.Query("SHOW TABLES")

	var table string

	for res.Next() {
		res.Scan(&table)
		fmt.Println(table)
	}
	db.Close()

	runner()

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func runner() {
	ctx, _ := context.WithCancel(context.Background())
	logger := logger.New(ctx, ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	iDatabase, err := mysql2.New(ctx, logger)
	if err != nil {
		logger.Error.Fatalln("can't connect to db")
	}

	iUseCases := usecases.New(ctx, iDatabase)

	var version dtoversion.Response

	_, err = http2.New(ctx, iUseCases, version, logger, true)
	if err != nil {
		logger.Error.Fatalln("err", err, "msg", "Error occurred for new http server", "func", "main")
	}
}
