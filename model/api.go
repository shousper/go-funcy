package model

import . "github.com/dave/jennifer/jen"

// Model type information for generation output
// Passed to fragments to create functions
type Model struct {
	Type        string
	IsInterface bool
	Map         MapModel
	Slice       FieldModel
	GroupBys    []FieldModel
}

func (m Model) TypeCode() Code {
	if m.IsInterface {
		return Id(m.Type)
	}
	return Op("*").Id(m.Type)
}

// FieldModel represents struct or interface field/method information used for generation
type FieldModel struct {
	Name     string
	Type     TypeModel
	Accessor string
}

// MapModel represents information for map fragment generation
type MapModel struct {
	Name string
	Key  *FieldModel
}

// TypeModel represents information about a type for generation
type TypeModel struct {
	Name      string
	Qualifier string
	Pointer   bool
}

func (t TypeModel) AsCode() Code {
	s := &Statement{}
	if t.Pointer {
		s = Op("*")
	}
	if t.Qualifier != "" {
		return s.Qual(t.Qualifier, t.Name)
	}
	return s.Id(t.Name)
}
