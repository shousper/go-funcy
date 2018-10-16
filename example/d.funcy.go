package example

import sub "github.com/shousper/go-funcy/example/sub"

// SliceOfD is a string slice of D
type SliceOfD []D

// GroupByE groups slice values by E()
func (s SliceOfD) GroupByE() map[string][]D {
	result := make(map[string][]D)
	for _, value := range s {
		result[value.E()] = append(result[value.E()], value)
	}
	return result
}

// GroupByApple groups slice values by Apple()
func (s SliceOfD) GroupByApple() map[*sub.Apple][]D {
	result := make(map[*sub.Apple][]D)
	for _, value := range s {
		result[value.Apple()] = append(result[value.Apple()], value)
	}
	return result
}

// AsMap maps slice values by E()
func (s SliceOfD) AsMap() map[string]D {
	result := make(map[string]D)
	for _, value := range s {
		result[value.E()] = value
	}
	return result
}

// MapOfD is a string map of D
type MapOfD map[string]D

// Keys returns a slice of the map keys
func (m MapOfD) Keys() (keys []string) {
	for key := range m {
		keys = append(keys, key)
	}
	return
}

// Values returns a slice of the map values
func (m MapOfD) Values() (values []D) {
	for _, value := range m {
		values = append(values, value)
	}
	return
}

// Select returns a map with only the given keys
func (m MapOfD) Select(keys ...string) MapOfD {
	result := make(MapOfD)
	for _, key := range keys {
		if value, exists := m[key]; exists {
			result[key] = value
		}
	}
	return result
}

// GroupByE groups map values by E
func (m MapOfD) GroupByE() map[string][]D {
	result := make(map[string][]D)
	for _, value := range m {
		result[value.E()] = append(result[value.E()], value)
	}
	return result
}

// GroupByApple groups map values by Apple
func (m MapOfD) GroupByApple() map[*sub.Apple][]D {
	result := make(map[*sub.Apple][]D)
	for _, value := range m {
		result[value.Apple()] = append(result[value.Apple()], value)
	}
	return result
}
