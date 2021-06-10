package query

import (
	"context"
	"fmt"
	"math"

	"github.com/georgysavva/scany/sqlscan"
	"github.com/masihtehrani/books/app/entities/structs"
)

// Exec execute select Query on Query field of Query struct and scan the rows in result
// its parameter result and return a Meta struct that contain data about pagination of Query.
func (q *Query) Exec(ctx context.Context, result interface{}) (structs.Meta, error) {
	err := q.prepareQuery()
	if err != nil {
		return structs.Meta{}, err
	}

	q.parseArgs(q.args)

	err = sqlscan.Select(ctx, q.db, result, q.Query, q.args...)
	if err != nil {
		return structs.Meta{}, fmt.Errorf("sqlscan.Select() >> %w", err)
	}

	err = q.db.QueryRow(q.metaQuery, q.args...).Scan(&q.meta.Total)
	if err != nil {
		return structs.Meta{}, fmt.Errorf("exec >> q.db.QueryRow() >> %w", err)
	}

	q.meta.LastPage = int(math.Ceil(float64(q.meta.Total) / float64(q.meta.PageSize)))

	return q.meta, nil
}
