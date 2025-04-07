package lister

import (
	"github.com/go-universal/cast"
)

// Lister provides a comprehensive set of methods to manage pagination,
// sorting, filtering, metadata, and response structures for data queries.
type Lister interface {
	// SetPage sets the current page number for pagination control.
	SetPage(page uint64) Lister

	// Page retrieves the current page number.
	Page() uint64

	// Pages returns the total number of pages based on the total record count and limit per page.
	Pages() uint64

	// SetLimit specifies the maximum number of items per page.
	SetLimit(limit uint) Lister

	// Limit retrieves the current limit of items per page.
	Limit() uint

	// AddSort appends a new sorting condition to the list of sorting criteria.
	// 'sort' defines the field to sort by, while 'order' specifies ascending or descending order.
	AddSort(sort string, order Order) Lister

	// Sort returns the current active sorting configuration.
	Sort() []Sort

	// SetSearch defines the search keyword or phrase for filtering results.
	SetSearch(search string) Lister

	// Search retrieves the current search keyword or phrase.
	Search() string

	// SetFilters assigns a map of key-value pairs to filter data.
	SetFilters(filters map[string]any) Lister

	// AddFilter assigns a single filter condition using a key-value pair.
	AddFilter(key string, value any) Lister

	// Filters returns the current set of applied filters.
	Filters() map[string]any

	// Filter retrieves the value of a specified filter.
	Filter(key string) any

	// HasFilter checks whether a filter exists for the specified key.
	HasFilter(key string) bool

	// CastFilter converts a filter value for a specified key into a caster type for flexible data handling.
	CastFilter(key string) cast.Caster

	// AddMeta assigns a meta-information value identified by a key.
	AddMeta(key string, value any) Lister

	// Meta retrieves the metadata value for the given key.
	Meta(key string) any

	// HasMeta checks whether metadata exists for the specified key.
	HasMeta(key string) bool

	// MetaData retrieves the complete metadata list as a map.
	MetaData() map[string]any

	// CastMeta converts a metadata value for a specified key into a caster type for flexible data handling.
	CastMeta(key string) cast.Caster

	// SetTotal sets the total number of available records.
	// This method must called after all sets.
	SetTotal(total uint64) Lister

	// Total retrieves the total number of available records.
	Total() uint64

	// From calculates the starting position of records for the current page.
	From() uint64

	// To calculates the ending position of records for the current page.
	To() uint64

	// SQLSortOrder generates an SQL-compatible string representing the ORDER BY and LIMIT conditions.
	SQLSortOrder() string

	// Response builds a standardized JSON response containing pagination details and metadata.
	Response() map[string]any

	// ResponseWithData generates a JSON response with additional data, combined with pagination and metadata details.
	ResponseWithData(data any) map[string]any
}
