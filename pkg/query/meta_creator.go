package query

import "fmt"

func (q *Query) metaCreator(pageSize, pageNumber int) string {
	// set meta data in Query struct.
	if pageSize > q.maxPageSize || pageSize < 1 {
		pageSize = q.maxPageSize
	}

	if pageNumber == 0 {
		pageNumber = 1
	}

	// create 'limit' and 'offset' of Query.
	limit := fmt.Sprintf(" LIMIT %v ", pageSize)
	offset := fmt.Sprintf(" OFFSET %v ", (pageNumber-1)*pageSize)

	q.meta.PageSize = pageSize
	q.meta.PageNumber = pageNumber

	return limit + offset
}

// create meta query create a select Total(*) query
// from created query to supplement meta data.
func (q *Query) createMetaQuery(slct, body, filter, groupBy, orderBy string) string {
	return fmt.Sprintf("SELECT COUNT(*) FROM (%s%s%s%s%s) aliase_name", slct, body, filter, groupBy, orderBy)
}
