package lister

import (
	"strconv"
	"strings"
)

// Sort represents a sorting field and its order (ascending or descending).
type Sort struct {
	Field string `json:"field"` // Field to sort by
	Order Order  `json:"order"` // Sorting order (ASC or DESC)
}

// SQLSortGenerator defines a function type for generating SQL ORDER BY clauses.
type SQLSortGenerator func(sorts []Sort, from uint64, limit uint) string

// mySQLSorter generates a MySQL-compatible ORDER BY clause with LIMIT and OFFSET.
func mySQLSorter(sorts []Sort, from uint64, limit uint) string {
	var sb strings.Builder
	sb.WriteString(" ORDER BY ")

	// Append sorting fields and their respective orders
	for i, sort := range sorts {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString("`" + sort.Field + "` ")
		sb.WriteString(sort.Order.String())
	}

	// Add LIMIT and OFFSET (MySQL uses LIMIT offset, count)
	sb.WriteString(" LIMIT " + strconv.FormatUint(max(from, 1)-1, 10))
	sb.WriteString(", " + strconv.FormatUint(uint64(limit), 10))
	return sb.String()
}

// postgreSQLSorter generates a PostgreSQL-compatible ORDER BY clause with LIMIT and OFFSET.
func postgreSQLSorter(sorts []Sort, from uint64, limit uint) string {
	var sb strings.Builder
	sb.WriteString(" ORDER BY ")

	// Append sorting fields and their respective orders
	for i, sort := range sorts {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(`"` + sort.Field + `" `)
		sb.WriteString(sort.Order.String())
	}

	// Add LIMIT and OFFSET (PostgreSQL uses LIMIT count OFFSET offset)
	sb.WriteString(" LIMIT " + strconv.FormatUint(uint64(limit), 10))
	sb.WriteString(" OFFSET " + strconv.FormatUint(from, 10))
	return sb.String()
}
