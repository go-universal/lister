package lister_test

import (
	"encoding/base64"
	"testing"

	"github.com/go-universal/lister"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParamConstructor(t *testing.T) {
	params := lister.ListerParams{
		Page:   3,
		Limit:  50,
		Search: "John Doe",
		Sorts: []lister.Sort{
			{Field: "name", Order: lister.Ascending},
			{Field: "id", Order: lister.Descending},
		},
		Filters: map[string]any{"author": "asc7DsX"},
	}

	lister := lister.NewFromParams(params)

	assert.Equal(t, uint64(3), lister.Page(), "Page() should return 3")
	assert.Equal(t, uint(50), lister.Limit(), "Limit() should return 50")
	assert.Equal(t, "John Doe", lister.Search(), `Search() should return "John Doe"`)
	assert.Equal(t, "asc7DsX", lister.CastFilter("author").StringSafe(""), `Filter() should return "asc7DsX"`)

	expected := ` ORDER BY "name" ASC, "id" DESC LIMIT 50 OFFSET 0`
	assert.Equal(t, expected, lister.SQLSortOrder(), "SQLSortOrder() should match the expected value")
}

func TestJsonConstructor(t *testing.T) {
	raw := `{
		"page": 2,
		"limit": 100,
		"search": "Jack ma",
		"filters": {
			"author": "asc7DsX"
		},
		"sorts": [
			{ "field": "name", "order": "desc" },
			{ "field": "id", "order": "asc" }
		]
	}`

	lister, err := lister.NewFromJson(raw)
	require.NoError(t, err, "NewFromJson() should not return an error")
	lister.SetTotal(1000)

	assert.Equal(t, uint64(2), lister.Page(), "Page() should return 2")
	assert.Equal(t, uint(100), lister.Limit(), "Limit() should return 100")
	assert.Equal(t, "Jack ma", lister.Search(), `Search() should return "Jack ma"`)
	assert.Equal(t, "asc7DsX", lister.CastFilter("author").StringSafe(""), `Filter() should return "asc7DsX"`)

	expected := ` ORDER BY "name" DESC, "id" ASC LIMIT 100 OFFSET 100`
	assert.Equal(t, expected, lister.SQLSortOrder(), "SQLSortOrder() should match the expected value")
}

func TestBase64Constructor(t *testing.T) {
	encoded := base64.URLEncoding.EncodeToString([]byte(
		`{
			"page": 2,
			"limit": 100,
			"search": "Jack ma",
			"filters": {
				"author": "asc7DsX"
			},
			"sorts": [
				{ "field": "name", "order": "desc" },
				{ "field": "id", "order": "asc" }
			]
		}`,
	))

	lister, err := lister.NewFromBase64Json(encoded)
	require.NoError(t, err, "NewFromBase64Json() should not return an error")
	lister.SetTotal(1000)

	assert.Equal(t, uint64(2), lister.Page(), "Page() should return 2")
	assert.Equal(t, uint(100), lister.Limit(), "Limit() should return 100")
	assert.Equal(t, "Jack ma", lister.Search(), `Search() should return "Jack ma"`)
	assert.Equal(t, "asc7DsX", lister.CastFilter("author").StringSafe(""), `Filter() should return "asc7DsX"`)

	expected := ` ORDER BY "name" DESC, "id" ASC LIMIT 100 OFFSET 100`
	assert.Equal(t, expected, lister.SQLSortOrder(), "SQLSortOrder() should match the expected value")
}
