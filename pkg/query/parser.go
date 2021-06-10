package query

import (
	"strings"
)

// ParseFilters take the params filter parse them,
// then create and return a Query string.
func (q *Query) parseFilters(filters map[string]map[string][]string) string {
	query := "WHERE "

	// range trough filters to create WHERE and WHERE IN clauses.
	for key, value := range filters {
		for k, v := range value {
			// look up Operations from operations map and use '=' as default.
			opr, ok := q.Operations[k]
			if !ok {
				opr = q.Operations["eq"]
			}

			query += q.where(Where{
				Field:     key,
				Values:    v,
				Operation: opr,
			})
		}

		query += "AND "
	}

	// trim redundant AND / WHERE at the end of Query if exist any.
	query = strings.TrimSuffix(query, "AND ")
	query = strings.TrimSuffix(query, "WHERE ")

	return query
}

// parseArgs parse the arguments and convert them from
// string  to the desired type.
func (q *Query) parseArgs(args []interface{}) {
	for i := range args {
		if args[i] == "true" {
			args[i] = true
		} else if args[i] == "false" {
			args[i] = false
		}
	}
}

// sanitizeFilters range trough filters and delete invalid filters.
func (q *Query) sanitizeFilters(selects []string,
	filters map[string]map[string][]string) map[string]map[string][]string {
	possibleFilters := make(map[string]string)

	for i := range selects {
		possibleFilters[dotSeparate(selects[i])] = selects[i]
	}

	sanitizedFilters := make(map[string]map[string][]string)

	for k, value := range filters {
		_, ok := possibleFilters[k]
		if ok {
			sanitizedFilters[k] = value
		}
	}

	for _, w := range q.wheres {
		m := make(map[string][]string)
		m[w.Operation] = w.Values
		sanitizedFilters[w.Field] = m
	}

	return sanitizedFilters
}

// dotSeparate check if string is in x.y format return y.
func dotSeparate(str string) string {
	if strings.Contains(str, ".") {
		return str[strings.Index(str, ".")+1:]
	}

	return str
}
