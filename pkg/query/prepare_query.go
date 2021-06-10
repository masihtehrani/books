package query

import (
	"fmt"
)

// prepareQuery create a select Query from filters that comes from http.

// nolint : funlen, nestif
func (q *Query) prepareQuery() error {
	// create select part of query.
	selct := q.createSelect()

	body := fmt.Sprintf("%s ", q.Body)

	// sanitize filters to remove redundant and invalid fields.
	filters := q.sanitizeFilters(q.Select, q.QueryFilters.Filter)

	// create 'where' part of the Query.
	where := q.parseFilters(filters)

	// create 'group by' part of the query.
	groupBy := q.groupBy()

	// create 'order by' part of the query.
	orderBy := q.orderBy(q.QueryFilters.Sort)

	// create limit/offset part of the query.
	limitOffset := q.metaCreator(q.QueryFilters.Page.Size, q.QueryFilters.Page.Number)

	// create meta query to supplement meta data.
	q.metaQuery = q.createMetaQuery(selct, body, where, groupBy, orderBy)

	// set Query filed of Query.
	q.Query = selct + body + where + groupBy + orderBy + limitOffset

	return nil
}

// createSelect creates select portion of Query.
func (q *Query) createSelect() string {
	query := "SELECT "

	// concatenate columns to select Query.
	for i := range q.Select {
		if i == 0 {
			query += q.Select[i]

			continue
		}

		query += fmt.Sprintf(", %s", q.Select[i])
	}

	return query + " "
}

// where takes a where struct and create desired 'where' or 'where in ' clause and return it.
func (q *Query) where(where Where) string {
	whereQuery := ""
	// whereLen is a number to consider if we have to use 'where in' or 'where' clause in query.
	const whereLen = 2

	if len(where.Values) == whereLen {
		whereQuery += fmt.Sprintf("%s IN(?,?) ", where.Field)
		q.args = append(q.args, where.Values[0], where.Values[1])
	} else {
		whereQuery += fmt.Sprintf("%s %s ? ", where.Field, where.Operation)
		q.args = append(q.args, where.Values[0])
	}

	return whereQuery
}

// orderBy takes a sort then create an 'ORDER BY' portion of a database select Query.

// nolint : nestif
func (q *Query) orderBy(sort map[string][]string) string {
	var orderBy string

	if len(sort) > 0 {
		orderBy += "ORDER BY "

		if asc, ok := sort["asc"]; ok {
			for i := range asc {
				if i == 0 {
					orderBy += fmt.Sprintf("%s ASC ", asc[i])

					continue
				}

				orderBy += fmt.Sprintf(", %s ASC ", asc[i])
			}
		}

		if desc, ok := sort["desc"]; ok {
			for i := range desc {
				if i == 0 {
					orderBy += fmt.Sprintf("%s DESC ", desc[i])

					continue
				}

				orderBy += fmt.Sprintf(", %s DESC ", desc[i])
			}
		}
	}

	return orderBy
}

func (q *Query) groupBy() string {
	var query string
	// concatenate GROUP BY part of Query if exist.
	if q.GroupBy != "" {
		query += fmt.Sprintf(" GROUP BY %s ", q.GroupBy)
	}

	return query
}
