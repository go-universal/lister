package lister

import "github.com/go-universal/cast"

// lister is a concrete implementation of the Lister interface.
type lister struct {
	option *option

	// Inputs
	page    uint64
	limit   uint
	search  string
	sorts   []Sort
	filters map[string]any

	// Pagination meta data
	total uint64
	from  uint64
	to    uint64
	pages uint64
	meta  map[string]any
}

// New creates a new instance of the Lister interface with customizable options.
func New(options ...Options) Lister {
	option := newOption()
	for _, opt := range options {
		opt(option)
	}

	return &lister{
		option:  option,
		sorts:   make([]Sort, 0),
		filters: make(map[string]any),
		meta:    make(map[string]any),
	}
}

func (l *lister) SetPage(page uint64) Lister {
	l.page = page
	return l
}

func (l *lister) Page() uint64 {
	return l.page
}

func (l *lister) Pages() uint64 {
	return l.pages
}

func (l *lister) SetLimit(limit uint) Lister {
	l.limit = l.option.validateLimit(limit)
	return l
}

func (l *lister) Limit() uint {
	if l.limit == 0 {
		return l.option.defaultLimit
	}
	return l.limit
}

func (l *lister) AddSort(sort string, order Order) Lister {
	if valid := l.option.validateSort(sort); valid == sort {
		l.sorts = append(l.sorts, Sort{Field: valid, Order: order})
	}
	return l
}

func (l *lister) Sort() []Sort {
	if len(l.sorts) > 0 {
		return l.sorts
	}
	return []Sort{{Field: l.option.defaultSort, Order: Ascending}}
}

func (l *lister) SetSearch(search string) Lister {
	l.search = search
	return l
}

func (l *lister) Search() string {
	return l.search
}

func (l *lister) SetFilters(filters map[string]any) Lister {
	if filters == nil {
		l.filters = make(map[string]any)
	} else {
		l.filters = filters
	}
	return l
}

func (l *lister) AddFilter(key string, value any) Lister {
	l.filters[key] = value
	return l
}

func (l *lister) Filters() map[string]any {
	return l.filters
}

func (l *lister) Filter(key string) any {
	return l.filters[key]
}

func (l *lister) HasFilter(key string) bool {
	_, exists := l.filters[key]
	return exists
}

func (l *lister) CastFilter(key string) cast.Caster {
	return cast.NewCaster(l.filters[key])
}

func (l *lister) AddMeta(key string, value any) Lister {
	l.meta[key] = value
	return l
}

func (l *lister) MetaData() map[string]any {
	return l.meta
}

func (l *lister) Meta(key string) any {
	return l.meta[key]
}

func (l *lister) HasMeta(key string) bool {
	_, exists := l.meta[key]
	return exists
}

func (l *lister) CastMeta(key string) cast.Caster {
	return cast.NewCaster(l.meta[key])
}

func (l *lister) SetTotal(total uint64) Lister {
	if total == 0 {
		l.page, l.total, l.pages, l.from, l.to = 0, 0, 0, 0, 0
	} else {
		limit := uint64(l.Limit())
		l.page = max(l.page, 1)
		l.total = total
		l.pages = (total + limit - 1) / limit
		l.page = min(l.page, l.pages)
		l.from = (max(l.page, 1) - 1) * limit
		l.to = min(l.from+limit, total)
	}
	return l
}

func (l *lister) Total() uint64 {
	return l.total
}

func (l *lister) From() uint64 {
	return l.from
}

func (l *lister) To() uint64 {
	return l.to
}

func (l *lister) SQLSortOrder() string {
	return l.option.sorter(l.Sort(), l.From(), l.Limit())
}

func (l *lister) Response() map[string]any {
	res := make(map[string]any)
	for k, v := range l.meta {
		res[k] = v
	}
	res["limit"] = l.Limit()
	res["sorts"] = l.Sort()
	res["search"] = l.Search()
	if l.total == 0 {
		res["page"], res["total"], res["from"], res["to"], res["pages"] = 0, 0, 0, 0, 0
	} else {
		res["page"] = l.Page()
		res["total"] = l.Total()
		res["from"] = l.From() + 1
		res["to"] = l.To()
		res["pages"] = l.Pages()
	}
	return res
}

func (l *lister) ResponseWithData(data any) map[string]any {
	response := l.Response()
	response["data"] = data
	return response
}
