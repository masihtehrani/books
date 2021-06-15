package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/masihtehrani/books/app/adapters/database/mysql"
	"github.com/masihtehrani/books/app/ports/dto/dtoversion"
	"github.com/masihtehrani/books/app/ports/transport/http"
	"github.com/masihtehrani/books/app/usecases"
	"github.com/masihtehrani/books/pkg/logger"
)

var versionTag, versionCommit, versionDate, serviceName string //nolint gochecknoglobals

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	logger := logger.New(ctx, ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	logger.Info.Println("msg", fmt.Sprintf("run %s with commit %s & commit date %s by version %s",
		serviceName, versionCommit, versionDate, versionTag),
		"func", "cmd", "when", "Bootstrapping project")

	iDatabase, err := mysql.New(ctx, logger)
	if err != nil {
		logger.Error.Fatalln("can't connect to db")
	}

	iUseCases := usecases.New(ctx, iDatabase)

	version := dtoversion.Response{
		VersionTag:    versionTag,
		VersionCommit: versionCommit,
		VersionDate:   versionDate,
		ServiceName:   serviceName,
	}

	httpServer, err := http.New(ctx, iUseCases, version, logger, false)
	if err != nil {
		logger.Error.Fatalln("err", err, "msg", "Error occurred for new http server", "func", "main")
	}

	go interruptHook(ctx, cancelFunc, iDatabase, httpServer, logger)

	os.Exit(1)
}
