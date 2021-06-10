package main

import (
	"context"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/masihtehrani/books/pkg/logger"
)

func interruptHook(ctx context.Context, cancelFunc context.CancelFunc, iDataStore, httpServer io.Closer,
	logger *logger.Logger) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	defer close(c)

	for {
		select {
		case <-c:
			stop(iDataStore, httpServer, logger)
			cancelFunc()

			return
		case <-ctx.Done():
			stop(iDataStore, httpServer, logger)
			cancelFunc()

			return
		}
	}
}

func stop(iDataStore, httpServer io.Closer, logger *logger.Logger) {
	if httpServer != nil {
		err := httpServer.Close()
		if err != nil {
			logger.Error.Println(
				"err", err,
				"msg", "Error occurred while closing http server.",
				"func", "stop",
				"when", "http server")
		}
	}

	if iDataStore != nil {
		err := iDataStore.Close()
		if err != nil {
			logger.Error.Println(
				"err", err,
				"msg", "Error occurred while closing iDataStore.",
				"func", "stop",
				"when", "datastore")
		}
	}
}
