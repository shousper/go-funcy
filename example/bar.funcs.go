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
func (s SliceOfBar) AsMap() MapOfBar {
	result := make(MapOfBar)
	for _, value := range s {
		result[value.ID] = value
	}
	return result
}

// MapOfBar is a string map of Bar
type MapOfBar map[string]*Bar

// Keys returns a slice of the map keys
func (m MapOfBar) Keys() (keys []string) {
	for key := range m {
		keys = append(keys, key)
	}
	return
}

// Values returns a slice of the map values
func (m MapOfBar) Values() (values []*Bar) {
	for _, value := range m {
		values = append(values, value)
	}
	return
}

// Select returns a map with only the given keys
func (m MapOfBar) Select(keys ...string) MapOfBar {
	result := make(MapOfBar)
	for _, key := range keys {
		if value, exists := m[key]; exists {
			result[key] = value
		}
	}
	return result
}

// GroupByID groups map values by ID
func (m MapOfBar) GroupByID() map[string][]*Bar {
	result := make(map[string][]*Bar)
	for _, value := range m {
		result[value.ID] = append(result[value.ID], value)
	}
	return result
}
