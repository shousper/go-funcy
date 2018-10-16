package sub

// SliceOfApple is a string slice of Apple
type SliceOfApple []*Apple

// GroupByBreed groups slice values by breed
func (s SliceOfApple) GroupByBreed() map[string][]*Apple {
	result := make(map[string][]*Apple)
	for _, value := range s {
		result[value.breed] = append(result[value.breed], value)
	}
	return result
}

// AsMap maps slice values by Color
func (s SliceOfApple) AsMap() map[[4]byte]*Apple {
	result := make(map[[4]byte]*Apple)
	for _, value := range s {
		result[value.Color] = value
	}
	return result
}

// MapOfApple is a string map of Apple
type MapOfApple map[[4]byte]*Apple

// Keys returns a slice of the map keys
func (m MapOfApple) Keys() (keys [][4]byte) {
	for key := range m {
		keys = append(keys, key)
	}
	return
}

// Values returns a slice of the map values
func (m MapOfApple) Values() (values []*Apple) {
	for _, value := range m {
		values = append(values, value)
	}
	return
}

// Select returns a map with only the given keys
func (m MapOfApple) Select(keys ...[4]byte) MapOfApple {
	result := make(MapOfApple)
	for _, key := range keys {
		if value, exists := m[key]; exists {
			result[key] = value
		}
	}
	return result
}

// GroupByBreed groups map values by Breed
func (m MapOfApple) GroupByBreed() map[string][]*Apple {
	result := make(map[string][]*Apple)
	for _, value := range m {
		result[value.breed] = append(result[value.breed], value)
	}
	return result
}
