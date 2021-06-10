package http

import (
	"context"

	"github.com/masihtehrani/books/app/ports/dto/dtoversion"
	"github.com/masihtehrani/books/pkg/server"
)

func versionHandler(_ context.Context, ver dtoversion.Response) server.MiddleFunc {
	return func(ctx context.Context, rawRequest server.RawRequest) (interface{}, error) {
		return dtoversion.Response{
			VersionTag:    ver.VersionTag,
			VersionCommit: ver.VersionCommit,
			VersionDate:   ver.VersionDate,
			ServiceName:   ver.ServiceName,
		}, nil
	}
}
