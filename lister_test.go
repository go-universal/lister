package lister_test

import (
	"testing"

	"github.com/go-universal/lister"
	"github.com/stretchr/testify/assert"
)

func TestEmptyLister(t *testing.T) {
	lister := lister.New().SetTotal(0)

	assert.Equal(t, uint64(0), lister.Page(), "Page() should be 0")
	assert.Equal(t, uint64(0), lister.Pages(), "Pages() should be 0")
	assert.Equal(t, uint64(0), lister.Total(), "Total() should be 0")
	assert.Equal(t, uint64(0), lister.From(), "From() should be 0")
	assert.Equal(t, uint64(0), lister.To(), "To() should be 0")

	expected := ` ORDER BY "id" ASC LIMIT 25 OFFSET 0`
	assert.Equal(t, expected, lister.SQLSortOrder(), "SQLSortOrder() should match the expected value")
}

func TestLister(t *testing.T) {
	lister := lister.New(
		lister.WithLimits(20, 10, 20, 30),
		lister.WithSorts("id", "id", "name", "mobile"),
	).
		AddSort("name", lister.ParseOrder(-1)).
		AddSort("mobile", "desc").
		SetLimit(100).
		SetPage(100).
		SetTotal(101)

	assert.Equal(t, uint64(6), lister.Page(), "Page() should be 6")
	assert.Equal(t, uint64(6), lister.Pages(), "Pages() should be 6")
	assert.Equal(t, uint(20), lister.Limit(), "Limit() should be 20")
	assert.Equal(t, uint64(101), lister.Total(), "Total() should be 101")
	assert.Equal(t, uint64(100), lister.From(), "From() should be 100")
	assert.Equal(t, uint64(101), lister.To(), "To() should be 101")

	expected := ` ORDER BY "name" DESC, "mobile" DESC LIMIT 20 OFFSET 100`
	assert.Equal(t, expected, lister.SQLSortOrder(), "SQLSortOrder() should match the expected value")
}
