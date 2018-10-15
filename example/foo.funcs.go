package example

// SliceOfFoo is a string slice of Foo
type SliceOfFoo []*Foo

// GroupByKey groups slice values by Key
func (s SliceOfFoo) GroupByKey() map[int][]*Foo {
	result := make(map[int][]*Foo)
	for _, value := range s {
		result[value.Key] = append(result[value.Key], value)
	}
	return result
}

// GroupByStringField groups slice values by StringField
func (s SliceOfFoo) GroupByStringField() map[string][]*Foo {
	result := make(map[string][]*Foo)
	for _, value := range s {
		result[value.StringField] = append(result[value.StringField], value)
	}
	return result
}

// GroupByEmbeddedField groups slice values by EmbeddedField
func (s SliceOfFoo) GroupByEmbeddedField() map[*Bar][]*Foo {
	result := make(map[*Bar][]*Foo)
	for _, value := range s {
		result[value.EmbeddedField] = append(result[value.EmbeddedField], value)
	}
	return result
}

// GroupByInterfaceField groups slice values by InterfaceField
func (s SliceOfFoo) GroupByInterfaceField() map[D][]*Foo {
	result := make(map[D][]*Foo)
	for _, value := range s {
		result[value.InterfaceField] = append(result[value.InterfaceField], value)
	}
	return result
}

// AsMap maps slice values by Key
func (s SliceOfFoo) AsMap() MapOfFoo {
	result := make(MapOfFoo)
	for _, value := range s {
		result[value.Key] = value
	}
	return result
}

// MapOfFoo is a string map of Foo
type MapOfFoo map[int]*Foo

// Keys returns a slice of the map keys
func (m MapOfFoo) Keys() (keys []int) {
	for key := range m {
		keys = append(keys, key)
	}
	return
}

// Values returns a slice of the map values
func (m MapOfFoo) Values() (values []*Foo) {
	for _, value := range m {
		values = append(values, value)
	}
	return
}

// Select returns a map with only the given keys
func (m MapOfFoo) Select(keys ...int) MapOfFoo {
	result := make(MapOfFoo)
	for _, key := range keys {
		if value, exists := m[key]; exists {
			result[key] = value
		}
	}
	return result
}

// GroupByKey groups map values by Key
func (m MapOfFoo) GroupByKey() map[int][]*Foo {
	result := make(map[int][]*Foo)
	for _, value := range m {
		result[value.Key] = append(result[value.Key], value)
	}
	return result
}

// GroupByStringField groups map values by StringField
func (m MapOfFoo) GroupByStringField() map[string][]*Foo {
	result := make(map[string][]*Foo)
	for _, value := range m {
		result[value.StringField] = append(result[value.StringField], value)
	}
	return result
}

// GroupByEmbeddedField groups map values by EmbeddedField
func (m MapOfFoo) GroupByEmbeddedField() map[*Bar][]*Foo {
	result := make(map[*Bar][]*Foo)
	for _, value := range m {
		result[value.EmbeddedField] = append(result[value.EmbeddedField], value)
	}
	return result
}

// GroupByInterfaceField groups map values by InterfaceField
func (m MapOfFoo) GroupByInterfaceField() map[D][]*Foo {
	result := make(map[D][]*Foo)
	for _, value := range m {
		result[value.InterfaceField] = append(result[value.InterfaceField], value)
	}
	return result
}
