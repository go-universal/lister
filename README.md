# Lister Package

![GitHub Tag](https://img.shields.io/github/v/tag/go-universal/lister?sort=semver&label=version)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-universal/lister.svg)](https://pkg.go.dev/github.com/go-universal/lister)
[![License](https://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/go-universal/lister/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-universal/lister)](https://goreportcard.com/report/github.com/go-universal/lister)
![Contributors](https://img.shields.io/github/contributors/go-universal/lister)
![Issues](https://img.shields.io/github/issues/go-universal/lister)

The `lister` package simplifies pagination, sorting, filtering, and metadata handling for data queries. It also provides tools to generate SQL-compatible clauses and build structured responses.

## Installation

```bash
go get github.com/go-universal/lister
```

## Features

- **Pagination**: Manage page numbers and limits and calculate pagination data.
- **Sorting**: Add multiple sort fields and orders.
- **Filtering**: Apply key-value filters.
- **Metadata**: Add and retrieve custom metadata.

## Usage

### Create a Lister

You can create a `Lister` instance using the following constructors:

- **Basic Constructor**:

```go
lister.New(options...)
```

- **From Parameters**:

```go
lister.NewFromParams(params, options...)
```

- **From JSON**:

```go
lister.NewFromJson(jsonString, options...)
```

- **From Base64 JSON**:

```go
lister.NewFromBase64Json(base64String, options...)
```

## Methods

### Pagination

- `SetPage(page uint64) Lister`: Set the current page number for pagination control.
- `Page() uint64`: Get the current page number.
- `SetLimit(limit uint) Lister`: Set the maximum number of items per page.
- `Limit() uint`: Get the current limit of items per page.
- `SetTotal(total uint64) Lister`: Set the total number of available records. **This method must be called to calculate pagination data.**
- `Total() uint64`: Get the total number of available records.
- `Pages() uint64`: Calculate the total number of pages based on the total record count and limit per page.
- `From() uint64`: Calculate the starting position of records for the current page.
- `To() uint64`: Calculate the ending position of records for the current page.

### Sorting

- `AddSort(sort string, order Order) Lister`: Add a sorting condition. The `sort` parameter specifies the field to sort by, and `order` specifies the sorting order (`asc` or `desc`).
- `Sort() []Sort`: Get the current active sorting configuration.

### Searching

- `SetSearch(search string) Lister`: Set a search keyword or phrase for filtering results.
- `Search() string`: Get the current search keyword or phrase.

### Filtering

- `SetFilters(filters map[string]any) Lister`: Set multiple filters using a map of key-value pairs.
- `AddFilter(key string, value any) Lister`: Add a single filter condition using a key-value pair.
- `Filters() map[string]any`: Get all applied filters as a map.
- `Filter(key string) any`: Get the value of a specific filter by its key.
- `HasFilter(key string) bool`: Check whether a filter exists for the specified key.
- `CastFilter(key string) cast.Caster`: Convert a filter value for a specified key into a caster type for flexible data handling.

### Metadata

- `AddMeta(key string, value any) Lister`: Add metadata identified by a key.
- `Meta(key string) any`: Get the metadata value for the specified key.
- `HasMeta(key string) bool`: Check whether metadata exists for the specified key.
- `MetaData() map[string]any`: Get all metadata as a map.
- `CastMeta(key string) cast.Caster`: Convert a metadata value for a specified key into a caster type for flexible data handling.

### SQL Generation

- `SQLSortOrder() string`: Generate an SQL-compatible string representing the `ORDER BY` and `LIMIT` conditions.

### Response

- `Response() map[string]any`: Build a standardized JSON response containing pagination details and metadata.
- `ResponseWithData(data any) map[string]any`: Build a JSON response with additional data, combined with pagination and metadata details.

## Examples

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/go-universal/lister"
)

func main() {
    l := lister.New().
        SetPage(1).
        SetLimit(10).
        AddSort("name", "asc").
        SetTotal(125)

    fmt.Println("Page:", l.Page())
    fmt.Println("Limit:", l.Limit())
    fmt.Println("Total Pages:", l.Pages())
    fmt.Println("SQL Sort Order:", l.SQLSortOrder())
}
```

### Using JSON Constructor

```go
package main

import (
    "fmt"
    "github.com/go-universal/lister"
)

func main() {
    raw := `{
        "page": 2,
        "limit": 20,
        "sorts": [
            { "field": "name", "order": "desc" }
        ]
    }`

    l, err := lister.NewFromJson(raw)
    if err != nil {
        panic(err)
    }
    l.SetTotal(125)

    fmt.Println("Page:", l.Page())
    fmt.Println("Limit:", l.Limit())
    fmt.Println("Total Pages:", l.Pages())
    fmt.Println("SQL Sort Order:", l.SQLSortOrder())
}
```

### Using Base64 JSON Constructor

```go
package main

import (
    "encoding/base64"
    "fmt"
    "github.com/go-universal/lister"
)

func main() {
    encoded := base64.URLEncoding.EncodeToString([]byte(`{
        "page": 1,
        "limit": 25,
        "sorts": [
            { "field": "name", "order": "asc" }
        ]
    }`))

    l, err := lister.NewFromBase64Json(encoded)
    if err != nil {
        panic(err)
    }
    l.SetTotal(125)

    fmt.Println("Page:", l.Page())
    fmt.Println("Limit:", l.Limit())
    fmt.Println("Total Pages:", l.Pages())
    fmt.Println("SQL Sort Order:", l.SQLSortOrder())
}
```

## License

This project is licensed under the [ISC License](LICENSE).
