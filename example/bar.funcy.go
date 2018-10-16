package example

// SliceOfBar is a string slice of Bar
type SliceOfBar []*Bar

// GroupByID groups slice values by ID
func (s SliceOfBar) GroupByID() map[string][]*Bar {
	result := make(map[string][]*Bar)
	for _, value := range s {
		result[value.ID] = append(result[value.ID], value)
	}
	return result
}

// AsMap maps slice values by ID
func (s SliceOfBar) AsMap() map[string]*Bar {
	result := make(map[string]*Bar)
	for _, value := range s {
		result[value.ID] = value
	}
	return result
}

// MapOfBar is a string map of Bar
type MapOfBar map[string]*Bar
