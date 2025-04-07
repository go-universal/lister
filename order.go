package lister

import (
	"fmt"
	"strings"
)

// Order represents the sorting order.
type Order string

const (
	// Ascending indicates ascending order.
	Ascending Order = "asc"

	// Descending indicates descending order.
	Descending Order = "desc"
)

// ParseOrder converts a value into an Order. Returns Descending for "-1" or "desc", otherwise Ascending.
func ParseOrder(v any) Order {
	switch strings.ToLower(fmt.Sprint(v)) {
	case "-1", "desc":
		return Descending
	default:
		return Ascending
	}
}

// String returns the uppercase string representation of the Order.
func (order Order) String() string {
	return strings.ToUpper(string(order))
}

// Numeric returns 1 for Ascending and -1 for Descending.
func (order Order) Numeric() int {
	if order == Ascending {
		return 1
	}
	return -1
}
