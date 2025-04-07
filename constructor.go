package lister

import (
	"encoding/base64"
	"encoding/json"
)

// ListerParams represents the input parameters for creating a Lister.
// It includes pagination, sorting, search, and filtering options.
type ListerParams struct {
	Page    uint64         `json:"page"`
	Limit   uint           `json:"limit"`
	Sorts   []Sort         `json:"sorts"`
	Search  string         `json:"search"`
	Filters map[string]any `json:"filters"`
}

// NewFromParams creates a new Lister instance using the provided parameters.
func NewFromParams(params ListerParams, options ...Options) Lister {
	lister := New(options...)
	lister.SetPage(params.Page)
	lister.SetLimit(params.Limit)
	for _, sort := range params.Sorts {
		lister.AddSort(sort.Field, sort.Order)
	}
	lister.SetSearch(params.Search)
	lister.SetFilters(params.Filters)
	return lister
}

// NewFromJson creates a new Lister instance from a JSON string.
func NewFromJson(data string, options ...Options) (Lister, error) {
	var params ListerParams
	if err := json.Unmarshal([]byte(data), &params); err != nil {
		return New(options...), err
	}
	return NewFromParams(params, options...), nil
}

// NewFromBase64Json creates a new Lister instance from a Base64-URL-encoded JSON string.
func NewFromBase64Json(data string, options ...Options) (Lister, error) {
	decodedData, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return New(options...), err
	}
	return NewFromJson(string(decodedData), options...)
}
