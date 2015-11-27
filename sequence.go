package sequence

import "sort"

type (
	// Elements is slice of items
	Elements []Element

	// Element data haveing id and its score, rate
	Element struct {
		items map[string]interface{}
	}

	// SortByFloat64 is sorter by float64 value
	SortByFloat64 struct {
		Elements
		key string
	}

	// FilterFunc is function to filter element
	FilterFunc func(Element) bool

	// MapFunc is function to map element
	MapFunc func(Element) Element
)

// Set sets a value for a given key
func (e *Element) Set(key string, value interface{}) {
	if e.items == nil {
		e.items = make(map[string]interface{})
	}
	e.items[key] = value
}

// Get gets a value for a given key
func (e *Element) Get(key string) (interface{}, bool) {
	if e.items == nil {
		e.items = make(map[string]interface{})
	}
	v, ok := e.items[key]
	return v, ok
}

// GetFloat64 gets float64 value for a given key
func (e *Element) GetFloat64(key string) float64 {
	val, _ := e.Get(key)
	assertedVal, ok := val.(float64)
	if !ok {
		assertedVal = 0
	}
	return assertedVal
}

// GetString gets string value for a given key
func (e *Element) GetString(key string) string {
	val, _ := e.Get(key)
	assertedVal, ok := val.(string)
	if !ok {
		assertedVal = ""
	}
	return assertedVal
}

// Len is implementation of Sort interface
func (e Elements) Len() int {
	return len(e)
}

// Swap is implementation of Sort interface
func (e Elements) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

// Less is implementation of Sort interface for Elements
func (s SortByFloat64) Less(i, j int) bool {
	return s.Elements[i].GetFloat64(s.key) < s.Elements[j].GetFloat64(s.key)
}

// SortByFloat64Desc is sort by score Desc
func (e Elements) SortByFloat64Desc(key string) {
	s := SortByFloat64{Elements: e, key: key}
	sort.Sort(sort.Reverse(s))
}

// Filter apply a given filter function for each element
func (e Elements) Filter(f FilterFunc) Elements {
	filtered := make(Elements, 0, e.Len())
	for i := range e {
		if f(e[i]) {
			filtered = append(filtered, e[i])
		}
	}
	return filtered
}

// Map concurrently apply a given map function for each element
func (e Elements) Map(f MapFunc) Elements {
	for i := range e {
		e[i] = f(e[i])
	}
	return e
}
