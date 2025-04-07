package lister

import (
	"slices"
	"strings"
)

// option holds configuration settings for query handling.
type option struct {
	defaultLimit uint             // Default query limit if none is specified
	defaultSort  string           // Default sort field if none is specified
	limits       []uint           // Allowed query limits
	sorts        []string         // Allowed sort fields
	sorter       SQLSortGenerator // Function to generate SQL sort clauses
}

// Options defines a function type for modifying lister options.
type Options func(*option)

// WithLimits configures the default limit and allowed limits for lister.
func WithLimits(def uint, valids ...uint) Options {
	// Remove invalid limits (e.g., zero values).
	valids = slices.DeleteFunc(valids, func(v uint) bool { return v == 0 })

	return func(o *option) {
		if def > 0 {
			o.defaultLimit = def
		}
		if len(valids) > 0 {
			o.limits = append([]uint{}, valids...)
		}
	}
}

// WithSorts configures the default sort field and allowed sort fields for lister.
func WithSorts(def string, valids ...string) Options {
	def = strings.TrimSpace(def)
	// Remove invalid sort fields (e.g., empty strings).
	valids = slices.DeleteFunc(valids, func(v string) bool { return strings.TrimSpace(v) == "" })

	return func(o *option) {
		if def != "" {
			o.defaultSort = def
		}
		o.sorts = append([]string{}, valids...)
	}
}

// WithDefaultLimit sets the default query limit.
func WithDefaultLimit(def uint) Options {
	return func(o *option) {
		if def > 0 {
			o.defaultLimit = def
		}
	}
}

// WithDefaultSort sets the default sort field.
func WithDefaultSort(def string) Options {
	def = strings.TrimSpace(def)
	return func(o *option) {
		if def != "" {
			o.defaultSort = def
		}
	}
}

// WithMySQLSorter configures the SQL sorter to use MySQL syntax.
func WithMySQLSorter() Options {
	return func(o *option) {
		o.sorter = mySQLSorter
	}
}

// WithPostgreSQLSorter configures the SQL sorter to use PostgreSQL syntax.
func WithPostgreSQLSorter() Options {
	return func(o *option) {
		o.sorter = postgreSQLSorter
	}
}

// newOption creates and initializes a new option instance with default settings.
func newOption() *option {
	return &option{
		defaultLimit: 25,
		defaultSort:  "id",
		limits:       []uint{25, 50, 100, 250},
		sorts:        []string{},
		sorter:       postgreSQLSorter,
	}
}

// validateLimit checks if the provided limit is valid; if not, it returns the default limit.
func (o *option) validateLimit(v uint) uint {
	if v > 0 && (len(o.limits) == 0 || slices.Contains(o.limits, v)) {
		return v
	}
	return o.defaultLimit
}

// validateSort checks if the provided sort field is valid; if not, it returns the default sort field.
func (o *option) validateSort(s string) string {
	s = strings.TrimSpace(s)
	if s != "" && (len(o.sorts) == 0 || slices.Contains(o.sorts, s)) {
		return s
	}
	return o.defaultSort
}
